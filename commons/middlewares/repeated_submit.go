package middlewares

import "github.com/gin-gonic/gin"

type Store interface {
	Set(key, value string) error
	Get(key string) (string, bool)
}

func RepeatedSubmit(s Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
