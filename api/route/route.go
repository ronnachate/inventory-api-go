package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter sets up the router.
func SetupRouter(db *gorm.DB, timeout time.Duration, gin *gin.Engine) {
	routerGroup := gin.Group("")

	NewUserRouter(timeout, db, routerGroup)
}
