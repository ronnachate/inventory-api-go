package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ronnachate/inventory-api-go/controllers"
)

// SetupRouter sets up the router.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
	}

	return r
}
