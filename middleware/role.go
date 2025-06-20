package middleware

import (
	// "flashbook/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RBAC(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied. Only " + requiredRole + " allowed.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
