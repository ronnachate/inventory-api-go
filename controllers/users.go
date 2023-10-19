package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers gets all existing users.
func GetUsers(c *gin.Context) {
	var users []string = []string{"John", "Jane", "Joe"}
	c.JSON(http.StatusOK, users)
}
