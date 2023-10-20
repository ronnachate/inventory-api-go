package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ronnachate/inventory-api-go/domain"
	"github.com/ronnachate/inventory-api-go/domain/mocks"
	"github.com/ronnachate/inventory-api-go/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByID(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	userID := uuid.UUID{}

	t.Run("success", func(t *testing.T) {

		mockUser := domain.User{
			ID: userID,
		}

		mockUserRepository.On("GetByID", mock.Anything, userID.String()).Return(mockUser, nil).Once()

		u := usecase.NewUserUsecase(mockUserRepository, time.Second*2)

		user, err := u.GetByID(context.Background(), userID.String())

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {

		//need to mock return empty user due to 'ret.Get(0).(domain.User)' error in generated file mocks/UserRepository.go
		mockUserRepository.On("GetByID", mock.Anything, userID.String()).Return(domain.User{}, errors.New("Unexpected")).Once()

		u := usecase.NewUserUsecase(mockUserRepository, time.Second*2)

		_, err := u.GetByID(context.Background(), userID.String())

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})
}
