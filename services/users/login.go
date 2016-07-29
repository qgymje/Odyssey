package users

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"
	"errors"
	"time"
)

var (
	// ErrLogin 登录失败
	ErrLogin = errors.New("登录失败, 手机号码或者密码错误")
)

// Login 用于登录操作
type Login struct {
	phone     string
	password  *Password
	userModel *models.User

	ensureDidFindUser bool
}

func NewLogin(data *forms.LoginForm) *Login {
	form := new(Login)
	form.phone = data.Phone
	form.password = NewPassword(data.Password)
	form.userModel = &models.User{}
	return form
}

func NewLoginByRawData(phone, password string) *Login {
	form := new(Login)
	form.phone = phone
	form.password = NewPassword(password)
	form.userModel = &models.User{}
	return form
}

func (s *Login) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.Login.Do error: ", err)
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

func (s *Login) findUser() (err error) {
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

func (s *Login) validPassword() error {
	if !s.ensureDidFindUser {
		if err := s.findUser(); err != nil {
			return err
		}
	}

	s.password.SetSalt(s.userModel.Salt)
	if !s.password.IsEncryptedSame(s.userModel.Password) {
		return ErrLogin
	}

	return nil
}

func (s *Login) updateToken() error {
	// generate token
	claims := map[string]interface{}{
		"id": s.userModel.ID,
	}
	token, err := NewToken().Generate(claims)
	if err != nil {
		return err
	}
	s.userModel.Token = models.NullString{String: token}

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

func (s *Login) UserInfo() *User {
	u := &User{
		ID:        s.userModel.ID,
		Phone:     s.userModel.Phone,
		Token:     s.userModel.Token.String,
		CreatedAt: s.userModel.CreatedAt,
	}
	return u
}
