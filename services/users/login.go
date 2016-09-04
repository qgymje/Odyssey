package users

import (
	"errors"
	"time"

	"github.com/qgymje/Odyssey/models"
	"github.com/qgymje/Odyssey/utils"
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
}

// LoginConfig 定义登录需要的数据
type LoginConfig struct {
	Phone    string
	Password string
}

func NewLogin(config *LoginConfig) *Login {
	l := new(Login)
	l.phone = config.Phone
	l.password = NewPassword(config.Password)
	l.userModel = &models.User{}
	return l
}

func NewLoginByRawData(phone, password string) *Login {
	l := new(Login)
	l.phone = phone
	l.password = NewPassword(password)
	l.userModel = &models.User{}
	return l
}

func (s *Login) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.Login.Do error: ", err)
		}
	}()

	if err = s.findUser(); err != nil {
		return
	}

	if err = s.validPassword(); err != nil {
		return
	}

	if err = s.updateToken(); err != nil {
		return
	}
	return
}

func (s *Login) findUser() (err error) {
	s.userModel, err = models.FindUserByPhone(s.phone)
	return
}

func (s *Login) validPassword() (err error) {
	s.password.SetSalt(s.userModel.Salt)
	if !s.password.IsEncryptedSame(s.userModel.Password) {
		return ErrLogin
	}
	return
}

func (s *Login) updateToken() (err error) {
	claims := map[string]interface{}{
		"id": s.userModel.ID,
	}
	token, err := NewToken().Generate(claims)
	if err != nil {
		return
	}
	err = s.userModel.UpdateToken(token)

	return
}

// User 输出结果
type UserInfo struct {
	ID        int64     `json:"id"`
	Phone     string    `json:"phone"`
	Nickname  string    `json:"nickname"`
	Token     string    `json:"token"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

// DefaultAvatar 生成根据nickname第一个字母的图片, 以及随机的背景颜色
func (u *UserInfo) DefaultAvatar() {

}

func (s *Login) UserInfo() *UserInfo {
	u := &UserInfo{
		ID:        s.userModel.ID,
		Phone:     s.userModel.Phone,
		Nickname:  s.userModel.Nickname.String,
		Avatar:    s.userModel.Avatar.String,
		Token:     s.userModel.Token.String,
		CreatedAt: s.userModel.CreatedAt,
	}
	return u
}
