package users

import (
	"errors"

	"github.com/qgymje/Odyssey/models"
	"github.com/qgymje/Odyssey/utils"
)

var (
	// ErrPhoneExists 手机号已存在错误
	ErrPhoneExists = errors.New("手机号码已经存在")
	// ErrSMSCode 验证码错误
	ErrSMSCode = errors.New("验证码错误")
)

// Register 注册对象
type Register struct {
	*Login
	smsCodeValidator *SMSCodeValidator
}

type RegisterConfig struct {
	*LoginConfig
	Code string
}

// NewRegister 生成一个注册用户对象
func NewRegister(config *RegisterConfig) *Register {
	s := new(Register)
	s.Login = NewLoginByRawData(config.Phone, config.Password)
	s.smsCodeValidator = NewSMSCodeValidator(config.Phone, config.Code)

	return s
}

// Do 做具体注册的操作
func (s *Register) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.Register.Do error: ", err)
		}
	}()

	if err = s.findPhone(); err != nil {
		return
	}

	if err = s.validSMSCode(); err != nil {
		return
	}

	if err = s.saveUser(); err != nil {
		return
	}

	if err = s.updateToken(); err != nil {
		return
	}

	if err = s.useSMSCode(); err != nil {
		return
	}

	return
}

func (s *Register) findPhone() (err error) {
	if models.IsPhoneRegisted(s.phone) {
		return ErrPhoneExists
	}
	return
}

func (s *Register) validSMSCode() (err error) {
	return s.smsCodeValidator.Valid()
}

func (s *Register) useSMSCode() (err error) {
	return s.smsCodeValidator.useCode()
}

// save 将数据保存到db
func (s *Register) saveUser() (err error) {
	if err = s.password.Valid(); err != nil {
		return
	}

	s.userModel.Phone = s.phone
	s.userModel.Salt = s.password.GenSalt()
	s.userModel.Password = s.password.GenPwd()

	return s.userModel.Create()
}
