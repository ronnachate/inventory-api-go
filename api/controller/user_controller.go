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

// GetUsers godoc
// @summary Get Users
// @description  Get user list with pagination
// @tags user
// @produce json
// @param page path int true "pagination page parameter"
// @param rows path int true "pagination rows parameter"
// @response 200 {object} []domain.User "OK"
// @response 400 {object} domain.ErrorResponse "Invalid page params"
// @response 400 {object} domain.ErrorResponse "Invalid rows params"
// @response 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router	/users [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	uintPage, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid page params"})
		infrastructure.Logger.Error().Msg(fmt.Sprintf("GetUsers error, Invalid page params with %s", page))
		return
	}

	var rows = c.DefaultQuery("rows", "10")
	uintRows, err := strconv.ParseUint(rows, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid rows params"})
		infrastructure.Logger.Error().Msg(fmt.Sprintf("GetUsers error, Invalid rows params with %s", rows))
		return
	}

	users, err := uc.UserUsecase.GetUsers(c, int(uintPage), int(uintRows))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Internal error"})
		infrastructure.Logger.Error().Msg(fmt.Sprintf("GetUsers error with %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserById godoc
// @summary Get User by id
// @description  Get user by id
// @tags user
// @Produce json
// @param id path int true "user id"
// @response 200 {object} domain.User "OK"
// @response 404 {object} domain.ErrorResponse "No user found"
// @Router	/users/{userId} [get]
func (uc *UserController) GetUserById(c *gin.Context) {
	user, err := uc.UserUsecase.GetByID(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "No user found"})
		infrastructure.Logger.Error().Msg(fmt.Sprintf("GetUserById error with %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}
