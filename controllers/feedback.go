package controllers

import "github.com/gin-gonic/gin"

type Feedback struct {
	Base
}

func (f *Feedback) Create(c *gin.Context) {

}

// 读取反馈数据,需要一个专门的auth_key
func (f *Feedback) Read(c *gin.Context) {

}
