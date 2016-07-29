package users

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

// ErrRequestInOneMinute 限制请求次数
var ErrRequestInOneMinute = errors.New("一分钟以后再尝试")

// ErrCodeNotExists 验证码还未生成
var ErrCodeNotExists = errors.New("验证码不存在")

var once sync.Once
var connection *amqp.Connection
var channel *amqp.Channel

func initSMS() {
	connection, channel = utils.GetAMQP()
}

// SMS 用于验证码操作服务
type SMS struct {
	phone        string
	code         string
	smscodeModel models.SMSCode
}

// NewSMS 用于生成一个验证码对象, 用于生成验证码
func NewSMS(form *forms.SMSCodeForm) *SMS {
	// make sure the mq service is started
	once.Do(initSMS)

	s := new(SMS)
	s.phone = form.Phone
	return s
}

func newSMSByRawData(phone string, code string) *SMS {
	s := new(SMS)
	s.phone = phone
	s.code = code
	return s
}

// Do 主要业务逻辑操作
func (s *SMS) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.SMS.Do error: ", err)
		}
	}()

	if models.IsPhoneExists(s.phone) {
		return ErrPhoneExists
	}

	if err = s.findSMSCode(); err != nil {
		log.Println(s.smscodeModel.CreatedAt)
		if time.Since(s.smscodeModel.CreatedAt) < 1*time.Minute {
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

func (s *SMS) findSMSCode() (err error) {
	s.smscodeModel, err = models.FindSMSCode(s.phone)
	return
}

func (s *SMS) generate() string {
	code := utils.RandomInt(100000, 1000000)
	s.code = fmt.Sprintf("%d", code)
	utils.GetLog().Debug("phone = %s ,sms code = %s", s.phone, s.code)
	return s.code
}

// GetCode 无论如何, 拿验证码
func (s *SMS) GetCode() string {
	return s.code
}

func (s *SMS) save() (err error) {
	s.smscodeModel = models.SMSCode{
		Phone:     s.phone,
		Code:      s.code,
		CreatedAt: time.Now(),
	}
	// save to db
	err = s.smscodeModel.Create()
	return
}

func (s *SMS) useCode() (err error) {
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
func (s *SMS) send() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.SMS.send error: ", err)
		}
	}()

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

// SMSValidator 用于验证sms cdoe
type SMSValidator struct {
	*SMS
}

// NewSMSValidator 用于用户提交验证码时, 判断验证码是否有效
func NewSMSValidator(phone string, code string) *SMSValidator {
	smsValidator := new(SMSValidator)
	smsValidator.SMS = newSMSByRawData(phone, code)
	return smsValidator
}

// Vaild 做验证
func (s *SMSValidator) Valid() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.SMSValidator.Valid error: ", err)
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
