package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronnachate/inventory-api-go/domain"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

// GetUsers gets all existing users.
func GetUsers(c *gin.Context) {
	var users []string = []string{"John", "Jane", "Joe"}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserById(c *gin.Context) {
	user, err := uc.UserUsecase.GetByID(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
