package controllers

import "github.com/gin-gonic/gin"

type RunComment struct {
	Base
}

// Index 返回一条跑步纪录的评论列表
func (rc *RunComment) Index(c *gin.Context) {

}

// Show 显示具体一条评论的信息
func (rc *RunComment) Show(c *gin.Context) {
}

// Comment 回复一个run
func (rc *RunComment) Comment(c *gin.Context) {
}

// Reply 回复一个comment
func (rc *RunComment) Reply(c *gin.Context) {
}
