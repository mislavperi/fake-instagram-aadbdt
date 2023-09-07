// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LogServiceUpload is an autogenerated mock type for the LogServiceUpload type
type LogServiceUpload struct {
	mock.Mock
}

// LogAction provides a mock function with given fields: userID, action
func (_m *LogServiceUpload) LogAction(userID int, action string) error {
	ret := _m.Called(userID, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(userID, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLogServiceUpload interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogServiceUpload creates a new instance of LogServiceUpload. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogServiceUpload(t mockConstructorTestingTNewLogServiceUpload) *LogServiceUpload {
	mock := &LogServiceUpload{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
