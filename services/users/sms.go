package users

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

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
	modelSmscode *models.SMSCode
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

// Valid 验证输入的数据是否有效
func (s *SMS) Valid() (err error) {
	//手机号是否已经存在?
	s.generate()
	if err = s.save(); err != nil {
		return
	}
	if err = s.send(); err != nil {
		return
	}
	return
}

func (s *SMS) GetCode() string {
	return s.code
}

// Generate 生成一个验证码
func (s *SMS) generate() string {
	code := utils.RandomInt(100000, 1000000)
	s.code = fmt.Sprintf("%d", code)
	utils.GetLog().Debug("phone = %s ,sms code = %s", s.phone, s.code)
	return s.code
}

// 保存验证码
func (s *SMS) save() (err error) {
	s.modelSmscode = &models.SMSCode{
		Phone:     s.phone,
		Code:      s.code,
		CreatedAt: time.Now(),
	}
	// save to db
	return nil
}

type smsContent struct {
	Phone string `json:"phone"` // 目标手机号

	Content string `json:"content"` // 发送的内容
}

func (sc smsContent) marshal(phone string, content string) (body []byte, err error) {
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
		log.Fatal(err)
	}

	var sc smsContent
	var body []byte

	if body, err = sc.marshal(s.phone, s.code); err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	// should add a callback function
	once.Do(confirm)
	return nil
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
	sms *SMS
}

// NewSMSValidator 用于用户提交验证码时, 判断验证码是否有效
func NewSMSValidator(phone string, code string) *SMSValidator {
	smsValidator := new(SMSValidator)
	smsValidator.sms = newSMSByRawData(phone, code)
	return smsValidator
}

// Vaild 做验证
func (s *SMSValidator) Valid(code string) error {
	return nil
}

// 判断这个phone号码是否已经请求过验证码了
func (s *SMSValidator) isRequestedCode() bool {
	return true
}
