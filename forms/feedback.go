package forms

import (
	"Odyssey/services/users"

	"github.com/gin-gonic/gin"
)

type FeedbackForm struct {
	UserID  int64
	Content string `form:"content" binding:"required"`

	*Base
}

func NewFeedbackForm(c *gin.Context, user *users.UserInfo) (*FeedbackForm, error) {
	form := &FeedbackForm{
		Base:   newBase(),
		UserID: user.ID,
	}

	if err := c.Bind(form); err != nil {
		form.Msg.formatBindError(err)
		return form, err
	}

	if err := form.Valid(); err != nil {
		return form, err
	}
	return form, nil
}

func (fb *FeedbackForm) Valid() error {
	return nil
}

type FeedbackReplyForm struct {
	FeedbackID int64
	Reply      string `form:"reply" binding:"required"`

	*Base
}

func NewFeedbackReplyForm(c *gin.Context, feedbackID int64) (*FeedbackReplyForm, error) {
	form := &FeedbackReplyForm{
		FeedbackID: feedbackID,
		Base:       newBase(),
	}

	if err := c.Bind(form); err != nil {
		form.Msg.formatBindError(err)
		return form, err
	}

	if err := form.Valid(); err != nil {
		return form, err
	}
	return form, nil

}

func (fr *FeedbackReplyForm) Valid() (err error) {
	if err = fr.validReply(); err != nil {
		return
	}
	return
}

func (fr *FeedbackReplyForm) validReply() (err error) {
	if fr.Reply == "" {
		fr.Msg.setError("reply", "回复为空")
		return
	}
	return
}
