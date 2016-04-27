package services

import (
	"Odyssey/forms"
	"Odyssey/models"
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
	*SignIn
	smsValidator *SMSValidator

	model_user *models.User
}

func NewSignUp(data *forms.SignUpForm) *SignUp {
	s := new(SignUp)

	s.SignIn = &SignIn{
		phone:    NewPhone(data.Phone),
		password: NewPassword(data.Password),
	}

	//s.smsValidator = NewSMSValidator()
	s.model_user = &models.User{}
	return s
}

// 将数据保存到db
func (s *SignUp) save() error {
	s.model_user.Phone = s.phone.PhoneNumber()
	s.model_user.Salt = s.password.GenSalt()
	s.model_user.Password = s.password.GenPwd()

	if err := s.model_user.Create(); err != nil {
		return err
	}

	// generate token
	claims := map[string]interface{}{
		"id": s.model_user.Id,
	}
	token, err := NewToken().Generate(claims)
	if err != nil {
		return err
	}
	s.model_user.Token = token

	where := map[string]interface{}{
		"id": s.model_user.Id,
	}
	update := map[string]interface{}{
		"token": token,
	}

	if err := s.model_user.Update(where, update); err != nil {
		return err
	}

	return nil
}

type User struct {
	Id    uint64
	Token string
}

func (s *SignUp) UserInfo() *User {
	u := &User{
		Id:    s.model_user.Id,
		Token: s.model_user.Token,
	}
	return u
}

func (s *SignUp) validSMSCode() error {
	return nil
}

func (s *SignUp) Do() error {
	// validate phone number is exists

	// save to db
	if err := s.save(); err != nil {
		return err
	}
	return nil
}
