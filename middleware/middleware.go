package middleware

import (
	"github.com/broboredo/locapp-api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SecurityToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Security-token")

		handler.Logger.Infof("HEADER token: %v", token)
		handler.Logger.Infof("ENV token: %v", os.Getenv("SECURITY_TOKEN"))

		if token != os.Getenv("SECURITY_TOKEN") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}
