package forms

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type RunCommentForm struct {
	RunID           int64  `form:"run_id" binding:"required"`
	Content         string `form:"content" binding:"required"`
	ParentCommentID int64  `form:"parent_comment_id"`
	UserID          int64

	*Base
}

func NewRunCommentForm(c *gin.Context, userID int64) (*RunCommentForm, error) {
	form := &RunCommentForm{
		Base:   newBase(),
		UserID: userID,
	}

	if err := c.Bind(form); err != nil {
		log.Printf("%v\n", c.Errors.Errors())
		form.Msg.formatBindError2(c.Errors)
		return form, err
	}

	if err := form.Valid(); err != nil {
		return form, err
	}
	return form, nil
}

func (c *RunCommentForm) Valid() (err error) {
	if c.RunID <= 0 {
		return errors.New("run_id 错误")
	}

	if len(strings.Trim(c.Content, " ")) == 0 {
		return errors.New("content不能为空白")
	}
	return
}
