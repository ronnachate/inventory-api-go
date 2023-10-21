package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronnachate/inventory-api-go/domain"
	"github.com/ronnachate/inventory-api-go/infrastructure"
)

const appkeyHeader = "application_key"

func ApplicationKeyMiddleware(config infrastructure.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		appkeyHeader := c.Request.Header.Get(appkeyHeader)
		if config.RunningEnv != "development" {
			if appkeyHeader != config.ApplicationKey {
				c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
				c.Abort()
				return
			}
		}
	}
}
