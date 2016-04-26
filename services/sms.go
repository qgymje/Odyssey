package services

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"
	"fmt"
	"time"
)

type SMS struct {
	phone         *Phone
	code          string
	model_smscode *models.SMSCode
}

func NewSMS(data *forms.SMSCodeForm) *SMS {
	s := new(SMS)
	s.phone = NewPhone(data.Phone)
	return s
}

func newSMSByRawData(phone string, code string) *SMS {
	s := new(SMS)
	s.phone = data.Phone
	s.code = data.code
	return s
}

func (s *SMS) Valid() error {

}

func randInt(min, max int) int {
	return min + utils.GetRand().Intn(max-min)
}

func (s *SMS) Generate() string {
	code := randInt(1000, 10000)
	s.code = fmt.Sprintf("%d", code)
	utils.GetLog().Debug("phone = %s ,sms code = %s", s.phone, s.code)
	return s.code
}

// 保存验证码
func (s *SMS) save() error {
	s.model_smscode = &models.SMSCode{
		Phone:     s.phone,
		Code:      s.code,
		CreatedAt: time.Now(),
	}
	return nil
}

// 用于难sms cdoe
type SMSValidator struct {
	*SMS
}

func NewSMSValidator(phone string, code string) *SMSValidator {
	smsValidator := new(SMSValidator)
	smsValidator.sms = newSMSByRawData(phone, code)
	return smsValidator
}

func (s *SMSValidator) Valid(code string) error {
	return nil
}

// 判断这个phone号码是否已经请求过验证码了
func (s *SMSValidator) IsRequestedCode() bool {
	return true
}
