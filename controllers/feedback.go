package controllers

import (
	"Odyssey/forms"
	"Odyssey/services/feedbacks"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Feedback struct {
	Base
}

func (f *Feedback) Create(c *gin.Context) {
	f.Authorization(c)

	form, err := forms.NewFeedbackForm(c, f.CurrentUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": form.Msg.ErrorMsg(),
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

func (f *Feedback) Show(c *gin.Context) {

}

func (f *Feedback) Reply(c *gin.Context) {
	idStr := c.Param("feedback_id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "feedback_id 为空",
			"meta":  f.Meta(c),
		})
		return
	}
	feedbackID, _ := strconv.Atoi(idStr)

	form, err := forms.NewFeedbackReplyForm(c, int64(feedbackID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": form.Msg.ErrorMsg(),
			"meta":  f.Meta(c),
		})
		return
	}

	fr := feedbacks.NewFeedbackReply(form)
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
