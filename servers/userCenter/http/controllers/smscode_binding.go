package controllers

import (
	"fmt"

	"github.com/qgymje/Odyssey/services/users"

	"github.com/gin-gonic/gin"
)

type SMSCodeBinding struct {
	Phone string `form:"phone" binding:"required"`

	config *users.SMSCodeConfig
	*BaseBinding
}

func NewSMSCodeBinding(c *gin.Context) (*SMSCodeBinding, error) {
	form := &SMSCodeBinding{
		BaseBinding: newBaseBinding(),
		config:      &users.SMSCodeConfig{},
	}

	if err := c.Bind(form); err != nil {
		form.Msg.formatBindError(err)
		return form, err
	}

	if err := form.Valid(); err != nil {
		return form, err
	}

	return form, nil
}

func (s *SMSCodeBinding) Valid() error {
	if err := s.validPhone(); err != nil {
		s.Msg.setError("phone", err.Error())
		return err
	}
	return nil
}

func (s *SMSCodeBinding) validPhone() error {
	if v := s.Validation.Mobile(s.Phone, "phone"); v.Ok {
		return nil
	}
	return fmt.Errorf("手机号码错误: %s", s.Phone)
}

func (s *SMSCodeBinding) Config() *users.SMSCodeConfig {
	s.config.Phone = s.Phone
	return s.config
}
