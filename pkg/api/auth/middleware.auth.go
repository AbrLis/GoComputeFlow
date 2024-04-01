package auth

import (
	"github.com/gin-gonic/gin"
)

func ensureAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
