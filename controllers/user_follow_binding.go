package controllers

import (
	"Odyssey/services/users/follows"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserFollowBinding struct {
	FromUserID int64
	ToUserID   int64 `form:"to_user_id" bind:"required"`

	config *follows.UserFollowConfig
	*BaseBinding
}

func NewUserFollowBinding(c *gin.Context, userID int64) (*UserFollowBinding, error) {
	form := &UserFollowBinding{
		FromUserID:  userID,
		BaseBinding: newBaseBinding(),
		config:      &follows.UserFollowConfig{},
	}

	if err := c.Bind(form); err != nil {
		form.Msg.formatBindError(c.Err())
		return form, err
	}

	if err := form.Valid(); err != nil {
		return form, err
	}
	return form, nil
}

func (f *UserFollowBinding) Valid() (err error) {
	if f.ToUserID <= 0 {
		return errors.New("to_usre_id 错误")
	}
	return
}

func (f *UserFollowBinding) Config() *follows.UserFollowConfig {
	f.config.FromUserID = f.FromUserID
	f.config.ToUserID = f.ToUserID
	return f.config
}
