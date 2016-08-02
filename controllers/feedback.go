package controllers

import (
	"Odyssey/forms"
	"Odyssey/services/feedbacks"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Feedback struct {
	Base
}

func (f *Feedback) Create(c *gin.Context) {
	form, err := forms.NewFeedbackForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": form.ErrorMsg(),
			"meta":  f.Meta(c),
		})
		return
	}

	fb := feedbacks.NewFeedback(form)
	if err := fb.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  f.Meta(c),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}

// 读取反馈数据,需要一个专门的auth_key
func (f *Feedback) Index(c *gin.Context) {
	fbs, err := feedbacks.Find()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  f.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, fbs)
}
