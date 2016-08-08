package controllers

import (
	"Odyssey/services/users/follows"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserFollow struct {
	Base

	form *UserFollowForm
}

func (f *UserFollow) before(c *gin.Context) error {
	f.Authorization(c)

	form, err := NewUserFollowForm(c, f.CurrentUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  f.Meta(c),
		})
		return err
	}
	f.form = form
	return nil
}

func (f *UserFollow) Follow(c *gin.Context) {
	f.before(c)

	follow := follows.NewUserFollow(f.form)
	if err := follow.Follow(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  f.Meta(c),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200})
}

func (f *UserFollow) Unfollow(c *gin.Context) {
	f.before(c)

	follow := follows.NewUserFollow(f.form)
	if err := follow.UnFollow(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  f.Meta(c),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200})
}

func (f *UserFollow) Followers(c *gin.Context) {

}
