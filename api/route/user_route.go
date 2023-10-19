package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ronnachate/inventory-api-go/api/controller"
	"github.com/ronnachate/inventory-api-go/repository"
	"github.com/ronnachate/inventory-api-go/usecase"
	"gorm.io/gorm"
)

func NewUserRouter(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
	}
	group.GET("/users", uc.GetUsers)
	group.GET("/users/:id", uc.GetUserById)
}
