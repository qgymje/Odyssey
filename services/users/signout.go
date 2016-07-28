package users

import (
	"Odyssey/models"
	"errors"
)

var (
	ErrSignOut = errors.New("退出失败")
)

// SignOut 登出操作
type SignOut struct {
	token     string
	userModel *models.User

	ensureDidFindUser bool
}

// NewSignOut varify token and compare the claims ?
func NewSignOut(token string) *SignOut {
	s := new(SignOut)
	s.token = token
	return s
}

func (s *SignOut) Do() error {
	if err := s.varifyToken(); err != nil {
		return err
	}

	if err := s.findUser(); err != nil {
		return err
	}

	if err := s.removeToken(); err != nil {
		return err
	}
	return nil
}

func (s *SignOut) varifyToken() error {
	t := NewToken()
	if ok, err := t.Verify(s.token); !ok {
		return err
	}
	return nil
}

func (s *SignOut) findUser() (err error) {
	where := map[string]interface{}{
		"token=?": s.token,
	}
	s.userModel, err = models.FindUser(where)
	if err != nil {
		return ErrSignOut
	}
	s.ensureDidFindUser = true
	return nil
}

func (s *SignOut) removeToken() error {
	where := map[string]interface{}{
		"id=?": s.userModel.ID,
	}
	update := map[string]interface{}{
		"token=?": "",
	}

	if err := s.userModel.Update(where, update); err != nil {
		return err
	}
	return nil
}
