// middleware/logging_middleware.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request details
		log.Printf("%s - %s %s\n", c.ClientIP(), c.Request.Method, c.Request.URL.Path)

		// Pass control to the next handler
		c.Next()
	}
}
