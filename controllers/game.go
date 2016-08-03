package controllers

import "github.com/gin-gonic/gin"

type Game struct {
	Base
}

// Index 显示所有比赛
func (g *Game) Index(c *gin.Context) {

}

// Show 显示一场比赛的具体信息
func (g *Game) Show(c *gin.Context) {

}

// Create 创建一场比赛
func (g *Game) Create(c *gin.Context) {

}

// Update 更新一场比赛的信息
func (g *Game) Update(c *gin.Context) {

}

// Destroy 销毁一场还未发布的比赛
func (g *Game) Destroy(c *gin.Context) {

}
