package forms

import (
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type SignInForm struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`

	valid *validation.Validation

	errmsg
}

func NewSignInForm(c *gin.Context) (*SignInForm, error) {
	form := &SignInForm{}
	form.valid = &validation.Validation{}
	form.errmsg = newErrmsg()

	if err := c.Bind(form); err != nil {
		return nil, err
	}

	if err := form.Valid(); err != nil {
		return nil, err
	}
	return form, nil
}

func (s *SignInForm) Valid() error {
	if err := s.validPhone(); err != nil {
		s.setError("phone", err.Error())
		return err
	}
	return nil
}

func (s *SignInForm) validPhone() error {
	if v := s.valid.Mobile(s.Phone, "phone"); v.Ok {
		return nil
	}

	return fmt.Errorf("手机号码错误: %s", s.Phone)
}
