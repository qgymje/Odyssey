package services

import (
	"Odyssey/forms"
	"Odyssey/models"
	"errors"
)

var (
	ErrSignIn  = errors.New("登录失败, 手机号码或者密码错误")
	ErrSignIn2 = errors.New("登录失败, 手机号码或密码错误")
)

type SignIn struct {
	phone     string
	password  *Password
	userModel *models.User

	ensureDidFindUser bool
}

func NewSignIn(data *forms.SignInForm) *SignIn {
	form := new(SignIn)
	form.phone = data.Phone
	form.password = NewPassword(data.Password)
	form.userModel = &models.User{}
	return form
}

func NewSignInByRawData(phone, password string) *SignIn {
	form := new(SignIn)
	form.phone = phone
	form.password = NewPassword(password)
	form.userModel = &models.User{}
	return form
}

func (s *SignIn) findUser() error {
	where := map[string]interface{}{
		"phone": s.phone,
	}
	us, err := models.FindUsers(where)
	if err != nil {
		return err
	}
	if len(us) > 0 {
		s.userModel = us[0]
		s.ensureDidFindUser = true
	} else {
		return ErrSignIn2
	}
	return nil
}

func (s *SignIn) validPassword() error {
	if !s.ensureDidFindUser {
		if err := s.findUser(); err != nil {
			return err
		}
	}

	s.password.SetSalt(s.userModel.Salt)
	if !s.password.IsEncryptedSame(s.userModel.Password) {
		return ErrSignIn
	}

	return nil
}

func (s *SignIn) Do() error {
	if err := s.validPassword(); err != nil {
		return err
	}

	if err := s.updateToken(); err != nil {
		return err
	}
	return nil
}

func (s *SignIn) updateToken() error {
	// generate token
	claims := map[string]interface{}{
		"id": s.userModel.Id,
	}
	token, err := NewToken().Generate(claims)
	if err != nil {
		return err
	}
	s.userModel.Token = token

	where := map[string]interface{}{
		"id": s.userModel.Id,
	}
	update := map[string]interface{}{
		"token": token,
	}

	if err := s.userModel.Update(where, update); err != nil {
		return err
	}
	return nil
}

type User struct {
	Id    uint64 `json:"id"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}

func (s *SignIn) UserInfo() *User {
	u := &User{
		Id:    s.userModel.Id,
		Phone: s.userModel.Phone,
		Token: s.userModel.Token,
	}
	return u
}
