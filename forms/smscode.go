package forms

import (
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type SMSCodeForm struct {
	Phone string `form:"phone" binding:"required"`

	valid *validation.Validation
	*errmsg
}

func NewSMSCodeForm(c *gin.Context) (*SMSCodeForm, error) {
	form := &SMSCodeForm{}

	form.valid = &validation.Validation{}
	form.errmsg = newErrmsg()

	if err := c.Bind(form); err != nil {
		form.formatBindError(err)
		return form, err
	}

	if err := form.Valid(); err != nil {
		return form, err
	}

	return form, nil
}

func (s *SMSCodeForm) Valid() error {
	if err := s.validPhone(); err != nil {
		s.setError("phone", err.Error())
		return err
	}
	return nil
}

func (s *SMSCodeForm) validPhone() error {
	if v := s.valid.Mobile(s.Phone, "phone"); v.Ok {
		return nil
	}
	return fmt.Errorf("手机号码错误: %s", s.Phone)
}
