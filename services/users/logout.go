package users

import (
	"Odyssey/models"
	"Odyssey/utils"
	"errors"
)

var (
	ErrLogout       = errors.New("退出失败")
	ErrTokenUnvalid = errors.New("token不正确")
	ErrToeknUpdate  = errors.New("token更新出错")
)

// Logout 登出操作
type Logout struct {
	token     string
	userModel *models.User

	ensureDidFindUser bool
}

// NewLogout varify token and compare the claims ?
func NewLogout(token string) *Logout {
	s := new(Logout)
	s.token = token
	return s
}

func (s *Logout) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.Logout.Do error: ", err)
		}
	}()

	if err = s.varifyToken(); err != nil {
		return ErrTokenUnvalid
	}

	if err = s.findUserByToken(); err != nil {
		return ErrTokenNotFound
	}

	if err = s.removeToken(); err != nil {
		return
	}
	return
}

func (s *Logout) varifyToken() error {
	t := NewToken()
	if ok, err := t.Verify(s.token); !ok {
		return err
	}
	return nil
}

func (s *Logout) findUserByToken() (err error) {
	s.userModel, err = models.FindUserByToken(s.token)
	if err != nil {
		return ErrLogout
	}
	s.ensureDidFindUser = true
	return nil
}

func (s *Logout) removeToken() error {
	return s.userModel.RemoveToken()
}
