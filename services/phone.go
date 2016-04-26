package services

import (
	"errors"
	"time"

	"github.com/astaxie/beego/validation"
)

var ErrSMSDurationTooShort = errors.New("请求短信验证码过于频繁")
var ErrSendSMS = errors.New("发送短信验证码失败")
var ErrSMSCodeNotExists = errors.New("无效的短信验证码")
var ErrSMSCodeUsed = errors.New("短信验证码被使用过")

const (
	SMS_REGISTER_TEMPLATE_ID = 313
	SMS_REQ_DURATION         = 1 * time.Minute //验证码请求间隔
	DEFAULT_SMS_CODE_LENGTH  = 6
)

type SMSer interface {
	Send(templateId int, phone string, vars map[string]string) error
}

type Phone struct {
	phone string
	valid *validation.Validation
	code  string
	sms   *SMS
}

func NewPhone(phone string) *Phone {
	return &Phone{
		phone: phone,
		sms:   &SMS{},
		valid: &validation.Validation{},
	}
}

func (p *Phone) PhoneNumber() string {
	return p.phone
}

func (p *Phone) IsValid() bool {
	if v := p.valid.Mobile(p.phone, "phone"); v.Ok {
		return true
	}
	return false
}

func (p *Phone) IsExists() bool {
	return false
}

func (p *Phone) ValidSMSCode(code string) error {
	return p.sms.Valid(code)
}

func (p *Phone) GenCode(length int) string {
	p.code = p.sms.Generate()
	return p.code
}

// 防止单一手机号码无限次数被请求
func (p *Phone) IsRequestedCode() bool {
	p.sms.phone = p.phone
	return p.sms.IsRequestedCode()
}
