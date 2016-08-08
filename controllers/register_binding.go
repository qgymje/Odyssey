package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	*LoginForm
	Code string `form:"code" binding:"required"`
}

func NewRegisterForm(c *gin.Context) (*RegisterForm, error) {
	form := &RegisterForm{}
	var err error

	form.LoginForm, err = NewLoginForm(c)
	if err != nil {
		form.formatBindError(err)
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

func (s *RegisterForm) Valid() error {
	if err := s.LoginForm.Valid(); err != nil {
		return nil
	}

	if err := s.validCode(); err != nil {
		s.setError("code", err.Error())
		return err
	}

	return nil
}

func (s *RegisterForm) validCode() error {
	v := s.valid.Numeric(s.Code, "code")
	if v.Ok {
		return nil
	}

	return fmt.Errorf("验证码错误: %s", s.Code)
}
