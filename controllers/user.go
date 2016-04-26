package controllers

import "github.com/gin-gonic/gin"

// data binding and validation
type User struct {
	Base
}

func (u *User) SignUp(c *gin.Context) {
	// create a user
}

func (u *User) SignIn(c *gin.Context) {
}

func (u *User) SignOut(c *gin.Context) {
}
