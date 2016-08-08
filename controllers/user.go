package controllers

import (
	"Odyssey/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Base
}

// SMSCode 获取验证码
func (u *User) SMSCode(c *gin.Context) {
	form, err := NewSMSCodeBinding(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	sms := users.NewSMS(form.Config())
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
	form, err := NewRegisterBinding(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	reg := users.NewRegister(form.Config())
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
	form, err := NewLoginBinding(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  u.Meta(c),
		})
		return
	}

	li := users.NewLogin(form.Config())
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

func (u *User) FoundPassword(c *gin.Context) {

}

func (u *User) ResetPassword(c *gin.Context) {

}
func (u *User) Profile(c *gin.Context) {

}

func (u *User) Around(c *gin.Context) {

}

func (u *User) Games(c *gin.Context) {

}

func (u *User) Friends(c *gin.Context) {

}

func (u *User) Groups(c *gin.Context) {

}
