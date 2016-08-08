package controllers

import (
	"Odyssey/services/runs"

	"github.com/gin-gonic/gin"
)

type RunBinding struct {
	UserID       int64   `form:"user_id" binding:"required"`
	Distance     float64 `form:"distance" binding:"required"`
	Duration     int     `form:"duration" binding:"required"`
	IsPublic     bool    `form:"is_public" binding:"required"`
	Comment      string  `form:"comment" binding:"required"`
	RunLocations string  `form:"locations" binding:"required"`

	*BaseBinding

	config *runs.RunConfig
}

func NewRunBinding(c *gin.Context) (*RunBinding, error) {
	form := &RunBinding{
		BaseBinding: newBaseBinding(),
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

func (rf *RunBinding) Valid() error {
	return nil
}

func (rf *RunBinding) Config() *runs.RunConfig {
	rf.config.UserID = rf.UserID
	rf.config.Distance = rf.Distance
	rf.config.Duration = rf.Duration
	rf.config.IsPublic = rf.IsPublic
	rf.config.Comment = rf.Comment
	rf.config.RunLocations = rf.RunLocations
	return rf.config
}
