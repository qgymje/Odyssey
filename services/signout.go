package services

import (
	"Odyssey/models"
	"errors"
)

var (
	ErrSignOut = errors.New("退出失败")
)

type SignOut struct {
	token     string
	userModel *models.User

	ensureDidFindUser bool
}

// varify token and compare the claims ?
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

	if err := s.updateToken(); err != nil {
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

func (s *SignOut) findUser() error {
	where := map[string]interface{}{
		"token": s.token,
	}
	us, err := models.FindUsers(where)
	if err != nil {
		return err
	}
	if len(us) > 0 {
		s.userModel = us[0]
		s.ensureDidFindUser = true
	} else {
		return ErrSignOut
	}
	return nil
}

func (s *SignOut) updateToken() error {
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
