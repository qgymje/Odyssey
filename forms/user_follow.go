package forms

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type UserFollowForm struct {
	FromUserID int64
	ToUserID   int64 `form:"to_user_id" bind:"required"`

	*Base
}

func NewUserFollowForm(c *gin.Context, userID int64) (*UserFollowForm, error) {
	form := &UserFollowForm{
		FromUserID: userID,
		Base:       newBase(),
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

func (f *UserFollowForm) Valid() (err error) {
	if f.ToUserID <= 0 {
		return errors.New("to_usre_id 错误")
	}
	return
}
