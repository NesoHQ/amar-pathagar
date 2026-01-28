package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-library/internal/dto"
	"github.com/yourusername/online-library/internal/services"
)

func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.Error("missing authorization header"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, dto.Error("invalid authorization header format"))
			c.Abort()
			return
		}

		userID, role, err := authService.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.Error("invalid or expired token"))
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("user_role", role)
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, dto.Error("admin access required"))
			c.Abort()
			return
		}
		c.Next()
	}
}
