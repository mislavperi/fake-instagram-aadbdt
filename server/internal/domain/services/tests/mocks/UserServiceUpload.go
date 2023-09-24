// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	mock "github.com/stretchr/testify/mock"
)

// UserServiceUpload is an autogenerated mock type for the UserServiceUpload type
type UserServiceUpload struct {
	mock.Mock
}

// GetUserInformation provides a mock function with given fields: id
func (_m *UserServiceUpload) GetUserInformation(id int) (*models.User, error) {
	ret := _m.Called(id)

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserServiceUpload interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserServiceUpload creates a new instance of UserServiceUpload. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserServiceUpload(t mockConstructorTestingTNewUserServiceUpload) *UserServiceUpload {
	mock := &UserServiceUpload{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
