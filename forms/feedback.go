package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type FeedbackForm struct {
	UserID  int64  `form:"user_id" binding:"required"`
	Content string `form:"content" binding:"required"`

	valid *validation.Validation

	*errmsg
}

func NewFeedbackForm(c *gin.Context) (*FeedbackForm, error) {
	form := &FeedbackForm{}
	form.valid = &validation.Validation{}
	form.errmsg = newErrmsg()

	if err := c.Bind(form); err != nil {
		form.formatBindError(err)
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

	Base
}

func NewFeedbackReplyForm(c *gin.Context, feedbackID int64) (*FeedbackReplyForm, error) {
	form := &FeedbackReplyForm{
		FeedbackID: feedbackID,
		Base: Base{
			Validation: &validation.Validation{},
			errmsg:     newErrmsg()},
	}

	if err := c.Bind(form); err != nil {
		form.formatBindError(err)
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
		fr.setError("reply", "回复为空")
		return
	}
	return
}
