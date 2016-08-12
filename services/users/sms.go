package users

import (
	"Odyssey/models"
	"Odyssey/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	errorTrace "github.com/pkg/errors"

	"github.com/streadway/amqp"
)

// ErrRequestInOneMinute 限制请求次数
var ErrRequestInOneMinute = errors.New("一分钟以后再尝试")

// ErrCodeNotExists 验证码还未生成
var ErrCodeNotExists = errors.New("验证码不存在")

var once sync.Once
var connection *amqp.Connection
var channel *amqp.Channel

func initSMSCode() {
	connection, channel = utils.GetAMQP()
}

// SMSCode 用于验证码操作服务
type SMSCode struct {
	phone        string
	code         string
	smscodeModel *models.SMSCode
}

// SMSCodeConfig 用于此服务的配置数据
// 并且移除对上层form层的依赖
type SMSCodeConfig struct {
	Phone string
}

// NewSMSCode 用于生成一个验证码对象, 用于生成验证码
func NewSMSCode(config *SMSCodeConfig) *SMSCode {
	// make sure the mq service is started
	once.Do(initSMSCode)

	s := new(SMSCode)
	s.phone = config.Phone
	s.smscodeModel = &models.SMSCode{}
	return s
}

func newSMSCodeByRawData(phone string, code string) *SMSCode {
	s := new(SMSCode)
	s.phone = phone
	s.code = code
	return s
}

// Do 主要业务逻辑操作
func (s *SMSCode) Do() (err error) {
	defer func() {
		if err != nil {
			err = errorTrace.Wrap(err, "services.users.SMSCode.Do error")
			utils.GetLog().Error("%+v", err)
		}
	}()

	if models.IsPhoneRegisted(s.phone) {
		return ErrPhoneExists
	}

	if err = s.findSMSCode(); err == nil {
		if s.smscodeModel.IsGeneratedInDuration(1 * time.Minute) {
			return ErrRequestInOneMinute
		}
	}

	_ = s.generate()

	if err = s.save(); err != nil {
		return
	}
	if err = s.send(); err != nil {
		return
	}

	return
}

func (s *SMSCode) findSMSCode() (err error) {
	s.smscodeModel, err = models.FindSMSCodeByPhone(s.phone)
	return
}

func (s *SMSCode) generate() string {
	code := utils.RandomInt(100000, 1000000)
	s.code = fmt.Sprintf("%d", code)
	utils.GetLog().Debug("phone = %s ,sms code = %s", s.phone, s.code)
	return s.code
}

// GetCode 无论如何, 拿验证码
func (s *SMSCode) GetCode() string {
	return s.code
}

func (s *SMSCode) save() (err error) {
	s.smscodeModel = &models.SMSCode{
		Phone:     s.phone,
		Code:      s.code,
		CreatedAt: time.Now(),
	}
	// save to db
	err = s.smscodeModel.Create()
	return
}

func (s *SMSCode) useCode() (err error) {
	err = s.smscodeModel.UseCode()
	return
}

type smsContent struct {
	Phone   string `json:"phone"`   // 目标手机号
	Content string `json:"content"` // 发送的内容
}

func (sc smsContent) toJSON(phone string, content string) (body []byte, err error) {
	sc = smsContent{
		Phone:   phone,
		Content: content,
	}

	body, err = json.Marshal(&sc)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Send 将数据发送到rabbitmq里, 然后由rabbitmq的worker来发送短信到用户手机
func (s *SMSCode) send() (err error) {
	var (
		exchange     = "smssender"
		routingKey   = "smssender"
		exchangeType = "direct"
	)

	if err = channel.ExchangeDeclare(
		exchange,     // exchange name
		exchangeType, // exchange type, 类型: direct, fanout, topic, 三种不同的类型, 详情看mq手册
		true,         // durable 指exchange是否持久化, 也就是保存到nmesia数据库中
		false,        // autoDelete 是exchnage是否自动删除
		false,
		false,
		nil,
	); err != nil {
		return
	}

	var sc smsContent
	var body []byte

	if body, err = sc.toJSON(s.phone, s.code); err != nil {
		return
	}

	err = channel.Publish(
		exchange,
		routingKey, // routing key is here
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "application/json",
			ContentEncoding: "",
			DeliveryMode:    amqp.Transient,
			Body:            body,
			Priority:        0,
		},
	)
	if err != nil {
		return
	}

	// should add a callback function
	once.Do(confirm)
	return
}

func confirm() {
	// 确保消息收到
	//func (me *Channel) Confirm(noWait bool) error
	if err := channel.Confirm(false); err != nil {
		log.Fatal(err)
	}

	//func (me *Channel) NotifyPublish(confirm chan Confirmation) chan Confirmation
	confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))
	go func() {
		for {
			if confirmed := <-confirms; confirmed.Ack {
				log.Printf("confirmed ack, tag: %d", confirmed.DeliveryTag) // tag should be increased!
			} else {
				log.Printf("failed ack: tag: %d", confirmed.DeliveryTag)
			}
		}
	}()
}

// SMSCodeValidator 用于验证sms cdoe
type SMSCodeValidator struct {
	*SMSCode
}

// NewSMSCodeValidator 用于用户提交验证码时, 判断验证码是否有效
func NewSMSCodeValidator(phone string, code string) *SMSCodeValidator {
	smsValidator := new(SMSCodeValidator)
	smsValidator.SMSCode = newSMSCodeByRawData(phone, code)
	return smsValidator
}

// Vaild 做验证
func (s *SMSCodeValidator) Valid() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.SMSCodeValidator.Valid error: ", err)
		}
	}()

	if err = s.findSMSCode(); err == nil {
		if s.smscodeModel.Code != s.code {
			return ErrCodeNotExists
		}

		if s.smscodeModel.IsUsed() {
			return errors.New("验证码已被使用")
		}
	} else {
		return ErrCodeNotExists
	}

	return
}
