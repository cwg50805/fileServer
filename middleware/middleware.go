package middleware

import (
	"github.com/gin-gonic/gin"
)

// MiddleWare function
func MiddleWare(c *gin.Context) {
	success := true
	if success {
		c.Next()
	}

	c.Abort()
}
