package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ronnachate/inventory-api-go/domain"
	infrastructure "github.com/ronnachate/inventory-api-go/infrastructure"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

// GetUsers gets all existing users.
func (uc *UserController) GetUsers(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid page params"})
		infrastructure.Logger.Error().Msg(fmt.Sprintf("GetUsers error, Invalid page params with %s", page))
		return
	}

	var rows = c.DefaultQuery("rows", "10")
	intRows, err := strconv.Atoi(rows)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid rows params"})
		infrastructure.Logger.Error().Msg(fmt.Sprintf("GetUsers error, Invalid rows params with %s", rows))
		return
	}

	users, err := uc.UserUsecase.GetUsers(c, intPage, intRows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Internal error"})
		infrastructure.Logger.Error().Msg(fmt.Sprintf("GetUsers error with %s", err.Error()))
		return
	}

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
