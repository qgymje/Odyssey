package controllers

import "github.com/gin-gonic/gin"

type Order struct {
	Base
}

// Create 表示报名一个比赛
func (r *Order) Create(c *gin.Context) {

}

func (r *Order) Pay(c *gin.Context) {

}

func (r *Order) PayCancel(c *gin.Context) {

}

func (r *Order) PayRefund(c *gin.Context) {

}
