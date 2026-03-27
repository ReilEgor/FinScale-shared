package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

const headerUserID = "X-User-ID"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader(headerUserID)
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: user id missing"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), UserIDKey, userID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
