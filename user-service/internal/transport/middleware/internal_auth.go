package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InternalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.GetHeader("X-User-ID")
		if userIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			c.Abort()
			return
		}

		c.Set("user_id", uint(userID))
		c.Next()
	}
}
