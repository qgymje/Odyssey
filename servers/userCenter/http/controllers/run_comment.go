package controllers

import (
	"net/http"

	"github.com/qgymje/Odyssey/services/runs/comments"

	"github.com/gin-gonic/gin"
)

type RunComment struct {
	Base
}

// Index 返回一条跑步纪录的评论列表
func (rc *RunComment) Index(c *gin.Context) {

}

// Show 显示具体一条评论的信息
func (rc *RunComment) Show(c *gin.Context) {
}

// Comment 评论/回复一个run
func (rc *RunComment) Comment(c *gin.Context) {
	rc.Authorization(c)

	bs, err := NewRunCommentBinding(c, rc.CurrentUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  rc.Meta(c),
		})
		return
	}

	comment := comments.NewRunComment(bs.Config())
	if err = comment.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"meta":  rc.Meta(c),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         200,
		"run_comment_id": comment.CommentID(),
	})
}
