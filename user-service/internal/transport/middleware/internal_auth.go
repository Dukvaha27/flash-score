package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
