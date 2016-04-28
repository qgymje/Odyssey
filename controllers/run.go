package controllers

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/services/runs"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Run struct {
	Base
}

func (r *Run) Create(c *gin.Context) {
	var userId uint64
	var err error
	if userId, err = r.parseUserId(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, err := forms.NewRunForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, form.ErrorMsg())
		return
	}
	form.UserId = userId

	rs := runs.NewRun(form)
	if err := rs.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rs.RunInfo())
}

func (r *Run) parseUserId(c *gin.Context) (uint64, error) {
	idStr := c.Param("user_id") //string
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("用户id解析错误")
	}
	return idUint, nil
}

func (r *Run) parseRunId(c *gin.Context) (uint64, error) {
	idStr := c.Param("run_id") //string
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("Run id解析错误")
	}
	return idUint, nil
}

func (r *Run) Read(c *gin.Context) {
	var userId uint64
	var runId uint64
	var err error
	if userId, err = r.parseUserId(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if runId, err = r.parseRunId(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result []*models.Run
	result, err := runs.Find(userId, runId)
	if err != nil {

	}

}

func (r *Run) ReadOne(c *gin.Context) {
}
