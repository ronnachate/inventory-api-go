package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ronnachate/inventory-api-go/api/controller"
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
