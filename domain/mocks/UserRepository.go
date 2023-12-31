// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ronnachate/inventory-api-go/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: c, id
func (_m *UserRepository) GetByID(c context.Context, id string) (domain.User, error) {
	ret := _m.Called(c, id)

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.User, error)); ok {
		return rf(c, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: c, page, rows
func (_m *UserRepository) GetUsers(c context.Context, page int, rows int) ([]domain.User, error) {
	ret := _m.Called(c, page, rows)

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]domain.User, error)); ok {
		return rf(c, page, rows)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []domain.User); ok {
		r0 = rf(c, page, rows)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(c, page, rows)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
