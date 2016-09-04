package controllers

import (
	"net/http"

	"github.com/qgymje/Odyssey/services/feedbacks"

	"github.com/gin-gonic/gin"
)

type Feedback struct {
	Base
}

func (f *Feedback) Create(c *gin.Context) {
	f.Authorization(c)

	binding, err := NewFeedbackBinding(c, f.CurrentUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": binding.Msg.Error(),
			"meta":  f.Meta(c),
		})
		return
	}

	fb := feedbacks.NewFeedback(binding.Config())
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

func (f *Feedback) Show(c *gin.Context) {

}

func (f *Feedback) Reply(c *gin.Context) {
	f.Authorization(c)

	bs, err := NewFeedbackReplyBinding(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bs.Msg.Error(),
			"meta":  f.Meta(c),
		})
		return
	}

	fr := feedbacks.NewFeedbackReply(bs.Config())
	if err := fr.Do(); err != nil {
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
