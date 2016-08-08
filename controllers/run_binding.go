package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type RunForm struct {
	UserID       int64   `form:"user_id" binding:"required"`
	Distance     float64 `form:"distance" binding:"required"`
	Duration     int     `form:"duration" binding:"required"`
	IsPublic     bool    `form:"is_public" binding:"required"`
	Comment      string  `form:"comment" binding:"required"`
	RunLocations string  `form:"locations" binding:"required"`

	valid *validation.Validation

	*errmsg
}

func NewRunForm(c *gin.Context, userID int64) (*RunForm, error) {
	form := &RunForm{
		UserID: userID,
		valid:  &validation.Validation{},
		errmsg: newErrmsg(),
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

func (rf *RunForm) Valid() error {
	return nil
}
