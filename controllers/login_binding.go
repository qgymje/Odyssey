package controllers

import (
	"Odyssey/services/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginBinding struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`

	config *users.LoginConfig
	*BaseBinding
}

func NewLoginBinding(c *gin.Context) (*LoginBinding, error) {
	bs := &LoginBinding{
		BaseBinding: newBaseBinding(),
		config:      &users.LoginConfig{},
	}

	if err := c.Bind(bs); err != nil {
		bs.Msg.formatBindError(err)
		return bs, err
	}

	if err := bs.Valid(); err != nil {
		return bs, err
	}
	return bs, nil
}

func (s *LoginBinding) Valid() error {
	if err := s.validPhone(); err != nil {
		s.Msg.setError("phone", err.Error())
		return err
	}
	return nil
}

func (s *LoginBinding) validPhone() error {
	if v := s.Validation.Mobile(s.Phone, "phone"); v.Ok {
		return nil
	}
	return fmt.Errorf("手机号码错误: %s", s.Phone)
}

func (s *LoginBinding) Config() *users.LoginConfig {
	s.config.Phone = s.Phone
	s.config.Password = s.Password
	return s.config
}
