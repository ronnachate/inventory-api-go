package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ronnachate/inventory-api-go/api/controller"
	"github.com/ronnachate/inventory-api-go/domain"
	"github.com/ronnachate/inventory-api-go/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setUserID(userID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("x-user-id", userID)
		c.Next()
	}
}

func TestGetByID(t *testing.T) {
	userID := uuid.UUID{}

	t.Run("success", func(t *testing.T) {
		mockUser := domain.User{
			ID: userID,
		}

		mockUserUsecase := new(mocks.UserUsecase)

		mockUserUsecase.On("GetByID", mock.Anything, userID.String()).Return(mockUser, nil)

		gin := gin.Default()

		rec := httptest.NewRecorder()

		uc := &controller.UserController{
			UserUsecase: mockUserUsecase,
		}

		//setup router
		gin.GET("/users/:id", uc.GetUserById)

		body, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String(), nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("notfound", func(t *testing.T) {
		mockUserUsecase := new(mocks.UserUsecase)

		//need to mock return empty user due to 'ret.Get(0).(domain.User)' error in generated file mocks/UserUsecase.go
		//code checking error not nil in controller
		mockUserUsecase.On("GetByID", mock.Anything, userID.String()).Return(domain.User{}, errors.New("Unexpected"))

		gin := gin.Default()

		rec := httptest.NewRecorder()

		uc := &controller.UserController{
			UserUsecase: mockUserUsecase,
		}

		//setup router
		gin.GET("/users/:id", uc.GetUserById)

		body, err := json.Marshal(domain.ErrorResponse{Message: "No user found"})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String(), nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockUserUsecase.AssertExpectations(t)
	})
}
