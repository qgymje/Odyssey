package forms

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SignUpForm struct {
	*SignInForm
	Code string `form:"code" binding:"required"`
}

func NewSignUpForm(c *gin.Context) (*SignUpForm, error) {
	form := &SignUpForm{}
	var err error

	form.SignInForm, err = NewSignInForm(c)
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

func (s *SignUpForm) Valid() error {
	if err := s.SignInForm.Valid(); err != nil {
		return nil
	}

	if err := s.validCode(); err != nil {
		s.setError("code", err.Error())
		return err
	}

	return nil
}

func (s *SignUpForm) validCode() error {
	v := s.valid.Numeric(s.Code, "code")
	if v.Ok {
		return nil
	} else {
		fmt.Println(v.Error)
	}

	return fmt.Errorf("验证码错误: %d", s.Code)
}
