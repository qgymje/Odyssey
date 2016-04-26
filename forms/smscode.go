package forms

import "github.com/gin-gonic/gin"

type SMSCodeForm struct {
	Phone string `form:"phone" binding:"required"`
}

func NewSMSCodeForm(c *gin.Context) (*SMSCodeForm, error) {
	form := &SMSCodeForm{}
	if err := c.Bind(form); err != nil {
		return nil, err
	}
	return form, nil
}
