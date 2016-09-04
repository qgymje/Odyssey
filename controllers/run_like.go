package controllers

import (
	"net/http"

	"github.com/qgymje/Odyssey/services/runs/likes"

	"github.com/gin-gonic/gin"
)

type RunLike struct {
	Base

	form *RunLikeBinding
}

func (l *RunLike) before(c *gin.Context) error {
	l.Authorization(c)

	form, err := NewRunLikeBinding(c, l.CurrentUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  l.Meta(c),
		})
		return err
	}
	l.form = form
	return nil
}

func (l *RunLike) Like(c *gin.Context) {
	l.before(c)

	like := likes.NewRunLike(l.form.Config())
	if err := like.Like(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  l.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200})
}

func (l *RunLike) Unlike(c *gin.Context) {
	l.before(c)

	like := likes.NewRunLike(l.form.Config())
	if err := like.Unlike(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  l.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200})
}
