package users

import (
	"Odyssey/models"
	"Odyssey/utils"
	"errors"
)

var (
	ErrLogout = errors.New("退出失败")
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
		return
	}

	if err = s.findUser(); err != nil {
		return
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

func (s *Logout) findUser() (err error) {
	where := map[string]interface{}{
		"token=?": s.token,
	}
	s.userModel, err = models.FindUser(where)
	if err != nil {
		return ErrLogout
	}
	s.ensureDidFindUser = true
	return nil
}

func (s *Logout) removeToken() error {
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
