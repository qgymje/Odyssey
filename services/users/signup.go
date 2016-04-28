package users

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"
	"errors"
	"fmt"
)

var (
	ErrPhoneExists = errors.New("手机号码已经存在")
	ErrSMSCode     = errors.New("验证码错误")
)

// 注册
type SignUp struct {
	*SignIn
	smsValidator *SMSValidator
}

func NewSignUp(data *forms.SignUpForm) *SignUp {
	s := new(SignUp)

	s.SignIn = NewSignInByRawData(data.Phone, data.Password)

	//s.smsValidator = NewSMSValidator()

	return s
}

// 将数据保存到db
func (s *SignUp) save() error {
	utils.GetLog().Debug("s = %s", utils.Sdump(s))
	fmt.Println("s.phone = ", s.phone)

	s.userModel.Phone = s.phone
	s.userModel.Salt = s.password.GenSalt()
	s.userModel.Password = s.password.GenPwd()

	if err := s.userModel.Create(); err != nil {
		return err
	}

	if err := s.updateToken(); err != nil {
		return err
	}
	return nil
}

func (s *SignUp) validSMSCode() error {
	return nil
}

func (s *SignUp) findPhone() error {
	where := map[string]interface{}{
		"phone": s.phone,
	}
	us, err := models.FindUsers(where)
	if err != nil {
		us = nil
		return err
	}

	if len(us) > 0 {
		return ErrPhoneExists
	}
	return nil
}

func (s *SignUp) Do() error {
	// validate phone number is exists
	if err := s.findPhone(); err != nil {
		return err
	}

	// validate sms code

	// save to db
	if err := s.save(); err != nil {
		return err
	}
	return nil
}
