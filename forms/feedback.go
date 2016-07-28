package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type FeedbackForm struct {
	UserID  uint64 `form:"user_id" binding:"required"`
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
