package controllers

type Feedback struct {
	Base
}

/*
func (f *Feedback) Create(c *gin.Context) {
	form, err := forms.NewFeedbackForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, form.ErrorMsg())
		return
	}

	fb := feedbacks.NewFeedback(form)
	if err := fb.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}

// 读取反馈数据,需要一个专门的auth_key
func (f *Feedback) Read(c *gin.Context) {

}
*/
