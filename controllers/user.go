package controllers

import (
	"Odyssey/forms"
	"Odyssey/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Base
}

// 获取验证码
func (u *User) SMSCode(c *gin.Context) {
	form, err := forms.NewSMSCodeForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// passing form to service
	sms := services.NewSMS(form)
	code := sms.Generate()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

// 注册
func (u *User) SignUp(c *gin.Context) {
	form, err := forms.NewSignUpForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	services.NewSignUp(form)
}

func (u *User) SignIn(c *gin.Context) {
}

func (u *User) SignOut(c *gin.Context) {
}
