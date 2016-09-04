package controllers

import (
	"fmt"
	"log"

	"github.com/qgymje/Odyssey/services/users"

	"github.com/gin-gonic/gin"
)

type RegisterBinding struct {
	*LoginBinding
	Code string `form:"code" binding:"required"`

	config *users.RegisterConfig
}

func NewRegisterBinding(c *gin.Context) (*RegisterBinding, error) {
	form := &RegisterBinding{
		config: &users.RegisterConfig{},
	}
	var err error

	form.LoginBinding, err = NewLoginBinding(c)
	log.Println(form.Config())
	if err != nil {
		form.Msg.formatBindError(err)
		return form, err
	}

	if err = c.Bind(form); err != nil {
		return form, err
	}

	if err = form.Valid(); err != nil {
		return form, err
	}

	return form, nil
}

func (s *RegisterBinding) Valid() error {
	if err := s.LoginBinding.Valid(); err != nil {
		return nil
	}

	if err := s.validCode(); err != nil {
		s.Msg.setError("code", err.Error())
		return err
	}

	return nil
}

func (s *RegisterBinding) validCode() error {
	v := s.Validation.Numeric(s.Code, "code")
	if v.Ok {
		return nil
	}

	return fmt.Errorf("验证码错误: %s", s.Code)
}

func (s *RegisterBinding) Config() *users.RegisterConfig {
	s.config.Phone = s.LoginBinding.Phone
	s.config.Password = s.LoginBinding.Password
	s.config.Code = s.Code
	log.Println(s.config)
	return s.config
}
