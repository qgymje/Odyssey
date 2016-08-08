package controllers

import (
	"Odyssey/models"
	"Odyssey/services/runs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Run struct {
	Base
}

// Create 接收一个跑步纪录请求
func (r *Run) Create(c *gin.Context) {
	r.Authorization(c)

	bs, err := NewRunBinding(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(c),
		})
		return
	}

	rs := runs.NewRun(bs.Config())
	if err := rs.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, rs.RunInfo())
}

// Index 显示一个用户的所有跑步纪录
func (r *Run) Index(c *gin.Context) {
	var userID int
	var err error
	if userID, err = r.parseUserID(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(c),
		})
		return
	}

	var result []*models.Run
	result, err = runs.Find(int64(userID), r.GetPageNum(c), r.GetPageSize(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

//Show 显示一条跑步纪录
func (r *Run) Show(c *gin.Context) {
	var userID int
	var runID int
	var err error

	if userID, err = r.parseUserID(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(c),
		})
		return
	}

	if runID, err = r.parseRunID(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(c),
		})
		return
	}

	run, err := runs.FindOne(int64(userID), int64(runID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, run)
}
