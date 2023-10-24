package route

import (
	"time"

	"github.com/gin-gonic/gin"
	docs "github.com/ronnachate/inventory-api-go/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// SetupRouter sets up the router.
func SetupRouter(db *gorm.DB, timeout time.Duration, gin *gin.Engine) {
	docs.SwaggerInfo.Title = "Inventory API"
	docs.SwaggerInfo.BasePath = "/"
	routerGroup := gin.Group("")

	NewUserRouter(timeout, db, routerGroup)

	routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
