package services

import (
	"Odyssey/forms"
)

type SignIn struct {
	phone    *Phone
	password *Password
}

func NewSignIn(data *forms.SignInForm) *SignIn {
	return new(SignIn)
}

// 注册
type SignUp struct {
	signIn *SignIn
	smsValidator *SMSValidator
}

func NewSignUp(data *forms.SignUpForm) *SignUp {
	s := new(SignUp)
	s.code :=
	return s
}

func (s *SignUp) validSMSCode() error {

}

func (s *SignUp) Do() error {
	return nil
}
