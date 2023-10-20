package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronnachate/inventory-api-go/domain"
	infrastructure "github.com/ronnachate/inventory-api-go/infrastructure"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

// GetUsers gets all existing users.
func (uc *UserController) GetUsers(c *gin.Context) {

	var users []string = []string{"John", "Jane", "Joe"}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserById(c *gin.Context) {
	user, err := uc.UserUsecase.GetByID(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "No user found"})
		infrastructure.Logger.Error().Msg(fmt.Sprintf("GetUserById error with %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}
