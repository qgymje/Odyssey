package controllers

import (
	"errors"

	"github.com/qgymje/Odyssey/services/runs/likes"

	"github.com/gin-gonic/gin"
)

type RunLikeBinding struct {
	RunID  int64 `form:"run_id" binding:"required"`
	UserID int64

	config *likes.RunLikeConfig
	*BaseBinding
}

func NewRunLikeBinding(c *gin.Context, userID int64) (*RunLikeBinding, error) {
	form := &RunLikeBinding{
		UserID:      userID,
		BaseBinding: newBaseBinding(),
		config:      &likes.RunLikeConfig{},
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

func (l *RunLikeBinding) Valid() (err error) {
	if l.RunID <= 0 {
		return errors.New("run_id 错误")
	}
	return
}

func (l *RunLikeBinding) Config() *likes.RunLikeConfig {
	l.config.RunID = l.RunID
	l.config.UserID = l.UserID
	return l.config
}
