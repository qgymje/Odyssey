package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/qgymje/Odyssey/services/users"

	"github.com/gin-gonic/gin"
)

var defaultPageSize = 20

type Base struct {
	CurrentUser *users.UserInfo
}

func (b *Base) Authorization(c *gin.Context) {
	var err error
	// 拿到token
	authParser := users.NewHeaderTokenParser(c.Request)
	if err = authParser.Parse(); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		c.Abort()
	}
	tokenString := authParser.Token()
	// jwt验证token
	token := users.NewToken()
	if ok, err := token.Verify(tokenString); !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		c.Abort()
	}
	// set session of current user
	if b.CurrentUser, err = token.GetUserInfo(); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		c.Abort()
	}
	log.Println("current user: ", b.CurrentUser)
}

// Meta 在返回错误时候, 带上额外的信息
func (b *Base) Meta(c *gin.Context) map[string]interface{} {
	return map[string]interface{}{
		"timestamp": time.Now(),
		"url":       "http://" + c.Request.Host + c.Request.URL.String(),
	}
}

func (b *Base) GetPageNum(c *gin.Context) (page int) {
	page, _ = strconv.Atoi(c.Param("page"))
	return
}

func (b *Base) GetPageSize(c *gin.Context) (num int) {
	num, err := strconv.Atoi(c.Param("page_num"))
	if err != nil {
		num = defaultPageSize
	}
	return
}

func (r *Base) parseUserID(c *gin.Context) (id int, err error) {
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

func (r *Base) parseRunID(c *gin.Context) (id int, err error) {
	idStr := c.Param("run_id") //string
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("Run id解析错误")
	}
	return
}
