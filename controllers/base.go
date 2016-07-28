package controllers

import (
	"Odyssey/services/users"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Base struct {
}

func (b *Base) Authorization(c *gin.Context) {
	var err error
	authParser := users.NewHeaderTokenParser(c.Request)
	if err = authParser.Parse(); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		c.Abort()
	}
	tokenString := authParser.Token()
	token := users.NewToken()
	if ok, err := token.Verify(tokenString); !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		c.Abort()
	}
	// set session of current user
}

// Meta 在返回错误时候, 带上额外的信息
func (b *Base) Meta() map[string]interface{} {
	return map[string]interface{}{
		"timestamp": time.Now(),
	}
}
