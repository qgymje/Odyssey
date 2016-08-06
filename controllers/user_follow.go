package models

import "github.com/gin-gonic/gin"

type UserFollow struct {
	Base
}

func (f *UserFollow) Follow(c *gin.Context) {

}

func (f *UserFollow) Unfollow(c *gin.Context) {

}

func (f *UserFollow) Followers(c *gin.Context) {

}
