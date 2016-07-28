package controllers

import (
	"Odyssey/forms"
	"errors"
	"net/http"
	"strconv"

	"Odyssey/services/runs"

	"github.com/gin-gonic/gin"
)

type Run struct {
	Base
}

// Create 接收一个跑步纪录请求
func (r *Run) Create(c *gin.Context) {
	r.Authorization(c)

	var userID int
	var err error
	if userID, err = r.parseUserID(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	form, err := forms.NewRunForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": form.ErrorMsg(),
			"meta":  r.Meta(),
		})
		return
	}
	form.UserID = userID

	rs := runs.NewRun(form)
	if err := rs.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(),
		})
		return
	}

	c.JSON(http.StatusOK, rs.RunInfo())
}

func (r *Run) parseUserID(c *gin.Context) (id int, err error) {
	var idStr string
	idStr = c.PostForm("user_id")
	if idStr == "" {
		idStr = c.Param("user_id") //string
	}
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("用户id解析错误")
	}
	return
}

func (r *Run) parseRunID(c *gin.Context) (id int, err error) {
	idStr := c.Param("run_id") //string
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("Run id解析错误")
	}
	return
}

func (r *Run) Read(c *gin.Context) {
}

func (r *Run) ReadOne(c *gin.Context) {
	/*
		var userID int
		var runID int
		var err error
		if userID, err = r.parseUserID(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if runID, err = r.parseRunID(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var result *models.Run
		result, err = runs.FindOne(userID, runID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
	*/
}
