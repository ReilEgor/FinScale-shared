package middleware

import (
	"context"
	sharedContextUtil "github.com/ReilEgor/FinScale-shared/pkg/contextutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

const headerUserID = "X-User-ID"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader(headerUserID)
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: user id missing"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), sharedContextUtil.UserIDKey, userID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
