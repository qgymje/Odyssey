package controllers

import (
	"Odyssey/forms"
	"Odyssey/services/runs"
	"errors"
	"net/http"
	"strconv"

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
			"meta":  r.Meta(c),
		})
		return
	}

	form, err := forms.NewRunForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": form.ErrorMsg(),
			"meta":  r.Meta(c),
		})
		return
	}
	form.UserID = int64(userID)

	rs := runs.NewRun(form)
	if err := rs.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  r.Meta(c),
		})
		return
	}

	c.JSON(http.StatusOK, rs.RunInfo())
}

/*
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
	result, err = runs.Find(userID, r.GetPageNum(c), r.GetPageSize(c))
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

}
*/

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
