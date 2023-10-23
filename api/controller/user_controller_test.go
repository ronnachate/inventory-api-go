package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ronnachate/inventory-api-go/api/controller"
	"github.com/ronnachate/inventory-api-go/domain"
	"github.com/ronnachate/inventory-api-go/domain/mocks"
	infrastructure "github.com/ronnachate/inventory-api-go/infrastructure"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInit(t *testing.T) {
	logger := zerolog.New(os.Stdout)
	infrastructure.Logger = &logger
}

func TestGetUsers(t *testing.T) {
	userID := uuid.UUID{}
	userID2 := uuid.UUID{}
	page := 1
	rows := 10

	mockUser := domain.User{
		ID:       userID,
		StatusID: 0,
		Status:   domain.UserStatus{},
	}
	mockUser2 := domain.User{
		ID:       userID2,
		StatusID: 0,
		Status:   domain.UserStatus{},
	}

	t.Run("success without param", func(t *testing.T) {

		mockUserUsecase := new(mocks.UserUsecase)

		mockUserUsecase.On("GetUsers", mock.Anything, page, rows).Return([]domain.User{mockUser, mockUser2}, nil)

		gin := gin.Default()

		rec := httptest.NewRecorder()

		uc := &controller.UserController{
			UserUsecase: mockUserUsecase,
		}

		//setup router
		gin.GET("/users", uc.GetUsers)

		body, err := json.Marshal([]domain.User{mockUser, mockUser2})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("success with param", func(t *testing.T) {

		mockUserUsecase := new(mocks.UserUsecase)

		mockUserUsecase.On("GetUsers", mock.Anything, page, 1).Return([]domain.User{mockUser}, nil)

		gin := gin.Default()

		rec := httptest.NewRecorder()

		uc := &controller.UserController{
			UserUsecase: mockUserUsecase,
		}

		//setup router
		gin.GET("/users", uc.GetUsers)

		body, err := json.Marshal([]domain.User{mockUser})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/users?rows=1", nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserUsecase := new(mocks.UserUsecase)

		//need to mock return empty user due to 'ret.Get(0).(domain.User)' error in generated file mocks/UserUsecase.go
		//code checking error not nil in controller
		mockUserUsecase.On("GetUsers", mock.Anything, mock.Anything, mock.Anything).Return([]domain.User{}, errors.New("Unexpected"))

		gin := gin.Default()

		rec := httptest.NewRecorder()

		uc := &controller.UserController{
			UserUsecase: mockUserUsecase,
		}

		//setup router
		gin.GET("/users", uc.GetUsers)

		body, err := json.Marshal(domain.ErrorResponse{Message: "Internal error"})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("invalid page params", func(t *testing.T) {
		mockUserUsecase := new(mocks.UserUsecase)

		gin := gin.Default()

		rec := httptest.NewRecorder()

		uc := &controller.UserController{
			UserUsecase: mockUserUsecase,
		}

		//setup router
		gin.GET("/users", uc.GetUsers)

		body, err := json.Marshal(domain.ErrorResponse{Message: "Invalid page params"})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/users?page=string", nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("invalid rows params", func(t *testing.T) {
		mockUserUsecase := new(mocks.UserUsecase)

		gin := gin.Default()

		rec := httptest.NewRecorder()

		uc := &controller.UserController{
			UserUsecase: mockUserUsecase,
		}

		//setup router
		gin.GET("/users", uc.GetUsers)

		body, err := json.Marshal(domain.ErrorResponse{Message: "Invalid rows params"})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/users?rows=string", nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockUserUsecase.AssertExpectations(t)
	})
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
