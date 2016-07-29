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

// SMSCode 获取验证码
func (u *User) SMSCode(c *gin.Context) {
	form, err := forms.NewSMSCodeForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": form.ErrorMsg(),
			"meta":  u.Meta(c),
		})
		return
	}

	sms := users.NewSMS(form)
	if err := sms.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	code := sms.GetCode()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

// SignUp 注册
func (u *User) Register(c *gin.Context) {
	form, err := forms.NewRegisterForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	reg := users.NewRegister(form)
	if err := reg.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, reg.UserInfo())
}

// Login 登录 action
func (u *User) Login(c *gin.Context) {
	form, err := forms.NewLoginForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	li := users.NewLogin(form)
	if err := li.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, li.UserInfo())
}

// Logout 退出
func (u *User) Logout(c *gin.Context) {
	p := users.NewHeaderTokenParser(c.Request)
	if err := p.Parse(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	token := p.Token()
	lo := users.NewLogout(token)
	if err := lo.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200})
}
