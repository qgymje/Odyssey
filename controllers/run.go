package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Run struct {
	Base
}

func (r *Run) Create(c *gin.Context) {

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
}

func (r *Run) ReadOne(c *gin.Context) {
}
