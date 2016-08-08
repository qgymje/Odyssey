package controllers

import (
	"Odyssey/services/runs/likes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RunLike struct {
	Base
}

func (l *RunLike) Like(c *gin.Context) {
	l.Authorization(c)

	runID, err := l.parseRunID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "run_id 错误",
			"meta":  l.Meta(c),
		})
		return
	}

	like := likes.NewRunLike(int64(runID), l.CurrentUser.ID)
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
	l.Authorization(c)

	runID, err := l.parseRunID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "run_id 错误",
			"meta":  l.Meta(c),
		})
		return
	}

	like := likes.NewRunLike(int64(runID), l.CurrentUser.ID)
	if err := like.Unlike(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  l.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200})
}
