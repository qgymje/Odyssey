package controllers

import (
	"errors"
	"strings"

	"github.com/qgymje/Odyssey/services/runs/comments"

	"github.com/gin-gonic/gin"
)

type RunCommentBinding struct {
	RunID           int64  `form:"run_id" binding:"required"`
	Content         string `form:"content" binding:"required"`
	ParentCommentID int64  `form:"parent_comment_id"`
	UserID          int64

	config *comments.RunCommentConfig

	*BaseBinding
}

func NewRunCommentBinding(c *gin.Context, userID int64) (*RunCommentBinding, error) {
	bs := &RunCommentBinding{
		BaseBinding: newBaseBinding(),
		UserID:      userID,
		config:      &comments.RunCommentConfig{},
	}

	if err := c.Bind(bs); err != nil {
		bs.Msg.formatBindError2(c.Errors)
		return bs, err
	}

	if err := bs.Valid(); err != nil {
		return bs, err
	}
	return bs, nil
}

func (c *RunCommentBinding) Valid() (err error) {
	if c.RunID <= 0 {
		return errors.New("run_id 错误")
	}

	if len(strings.Trim(c.Content, " ")) == 0 {
		return errors.New("content不能为空白")
	}
	return
}

func (c *RunCommentBinding) Config() *comments.RunCommentConfig {
	c.config.RunID = c.RunID
	c.config.Content = c.Content
	c.config.ParentCommentID = c.ParentCommentID
	c.config.UserID = c.UserID
	return c.config
}
