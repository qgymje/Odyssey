package middlewares

import "github.com/gin-gonic/gin"

// user 1
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzI3MjkzNTcsImlkIjoxfQ.HSs7hZ89fwst05owcXLXenZGpZwWUm0_K6BhqxC2Frw"

func FakedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Add("Authorization", "bearer "+token)
		c.Next()
	}
}
