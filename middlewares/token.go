package middlewares

import "github.com/gin-gonic/gin"

// 拿出token
func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断url里是否有token字段
		// 判断header里是否有Authrozation
		// before request
		c.Set("token", "12345")
		c.Next()
	}
}
