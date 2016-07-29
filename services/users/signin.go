package users

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"
	"errors"
	"time"
)

var (
	// ErrSignIn 登录失败
	ErrSignIn = errors.New("登录失败, 手机号码或者密码错误")
)

// SignIn 用于登录操作
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

func (s *SignIn) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.SignIn.Do error: ", err)
		}
	}()

	if err = s.validPassword(); err != nil {
		return
	}

	if err = s.updateToken(); err != nil {
		return
	}
	return
}

func (s *SignIn) findUser() (err error) {
	where := map[string]interface{}{
		"phone=?": s.phone,
	}
	s.userModel, err = models.FindUser(where)
	if err != nil {
		return err
	}
	s.ensureDidFindUser = true
	return
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

func (s *SignIn) updateToken() error {
	// generate token
	claims := map[string]interface{}{
		"id": s.userModel.ID,
	}
	token, err := NewToken().Generate(claims)
	if err != nil {
		return err
	}
	s.userModel.Token = token

	where := map[string]interface{}{
		"id=?": s.userModel.ID,
	}
	update := map[string]interface{}{
		"token=?": token,
	}

	if err := s.userModel.Update(where, update); err != nil {
		return err
	}
	return nil
}

// User 输出结果
type User struct {
	ID        int       `json:"id"`
	Phone     string    `json:"phone"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *SignIn) UserInfo() *User {
	u := &User{
		ID:        s.userModel.ID,
		Phone:     s.userModel.Phone,
		Token:     s.userModel.Token,
		CreatedAt: s.userModel.CreatedAt,
	}
	return u
}
