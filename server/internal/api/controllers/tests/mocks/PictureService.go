// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	bytes "bytes"

	mock "github.com/stretchr/testify/mock"

	models "github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"

	multipart "mime/multipart"
)

// PictureService is an autogenerated mock type for the PictureService type
type PictureService struct {
	mock.Mock
}

// GetEditedImage provides a mock function with given fields: imageID, height, width, format, sepia, blur
func (_m *PictureService) GetEditedImage(imageID int, height int32, width int32, format string, sepia float32, blur float32) (*bytes.Buffer, error) {
	ret := _m.Called(imageID, height, width, format, sepia, blur)

	var r0 *bytes.Buffer
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int32, int32, string, float32, float32) (*bytes.Buffer, error)); ok {
		return rf(imageID, height, width, format, sepia, blur)
	}
	if rf, ok := ret.Get(0).(func(int, int32, int32, string, float32, float32) *bytes.Buffer); ok {
		r0 = rf(imageID, height, width, format, sepia, blur)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bytes.Buffer)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int32, int32, string, float32, float32) error); ok {
		r1 = rf(imageID, height, width, format, sepia, blur)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImageByID provides a mock function with given fields: id
func (_m *PictureService) GetImageByID(id int) (*models.Picture, error) {
	ret := _m.Called(id)

	var r0 *models.Picture
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Picture, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Picture); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Picture)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImages provides a mock function with given fields: filter
func (_m *PictureService) GetImages(filter models.Filter) ([]models.Picture, error) {
	ret := _m.Called(filter)

	var r0 []models.Picture
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Filter) ([]models.Picture, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(models.Filter) []models.Picture); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Picture)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Filter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserImages provides a mock function with given fields: userID
func (_m *PictureService) GetUserImages(userID int) ([]models.Picture, error) {
	ret := _m.Called(userID)

	var r0 []models.Picture
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]models.Picture, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) []models.Picture); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Picture)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateImageInformation provides a mock function with given fields: imageID, description, hashtags, userID
func (_m *PictureService) UpdateImageInformation(imageID int, description string, hashtags []string, userID int) error {
	ret := _m.Called(imageID, description, hashtags, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, []string, int) error); ok {
		r0 = rf(imageID, description, hashtags, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadImage provides a mock function with given fields: file, title, description, hashtags, userID, height, width, fileExt
func (_m *PictureService) UploadImage(file multipart.File, title string, description string, hashtags []string, userID int, height string, width string, fileExt string) error {
	ret := _m.Called(file, title, description, hashtags, userID, height, width, fileExt)

	var r0 error
	if rf, ok := ret.Get(0).(func(multipart.File, string, string, []string, int, string, string, string) error); ok {
		r0 = rf(file, title, description, hashtags, userID, height, width, fileExt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPictureService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPictureService creates a new instance of PictureService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPictureService(t mockConstructorTestingTNewPictureService) *PictureService {
	mock := &PictureService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
