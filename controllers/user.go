package controllers

import (
	"Odyssey/forms"
	"Odyssey/services/users"
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
		c.JSON(http.StatusBadRequest, form.ErrorMsg())
		return
	}

	sms := users.NewSMS(form)
	if err := sms.Valid(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	code := sms.GetCode()
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

	su := users.NewSignUp(form)
	if err := su.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, su.UserInfo())
}

func (u *User) SignIn(c *gin.Context) {
	form, err := forms.NewSignInForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	si := users.NewSignIn(form)
	if err := si.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, si.UserInfo())
}

func (u *User) SignOut(c *gin.Context) {
	p := users.NewHeaderTokenParser(c.Request)
	if err := p.Parse(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := p.Token()
	so := users.NewSignOut(token)
	if err := so.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200})
}

func (u *User) DeleteAccount(c *gin.Context) {

}
