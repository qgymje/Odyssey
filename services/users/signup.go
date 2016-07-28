package users

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"
	"errors"
)

var (
	// ErrPhoneExists 手机号已存在错误
	ErrPhoneExists = errors.New("手机号码已经存在")
	// ErrSMSCode 验证码错误
	ErrSMSCode = errors.New("验证码错误")
)

// SignUp 注册对象
type SignUp struct {
	*SignIn
	smsValidator *SMSValidator
}

// NewSignUp 生成一个注册用户对象
func NewSignUp(data *forms.SignUpForm) *SignUp {
	s := new(SignUp)
	s.SignIn = NewSignInByRawData(data.Phone, data.Password)
	s.smsValidator = NewSMSValidator(data.Phone, data.Code)

	return s
}

// Do 做具体注册的操作
func (s *SignUp) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.SignUp.Do error: ", err)
		}
	}()

	if err = s.findPhone(); err != nil {
		return
	}

	if err = s.smsValidator.Valid(); err != nil {
		return
	}
	if err = s.save(); err != nil {
		return
	}

	return
}

func (s *SignUp) findPhone() error {
	if models.IsPhoneExists(s.phone) {
		return ErrPhoneExists
	}
	return nil
}

// save 将数据保存到db
func (s *SignUp) save() (err error) {
	s.userModel.Phone = s.phone
	s.userModel.Salt = s.password.GenSalt()
	s.userModel.Password = s.password.GenPwd()

	if err = s.userModel.Create(); err != nil {
		return
	}

	if err = s.updateToken(); err != nil {
		return
	}

	if err = s.smsValidator.useCode(); err != nil {
		return
	}
	return
}
