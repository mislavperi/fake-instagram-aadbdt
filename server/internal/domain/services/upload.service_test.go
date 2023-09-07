package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services/tests/factories"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services/tests/mocks"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"github.com/stretchr/testify/mock"
)

func TestGetConsumption_TotalConsumptionExceedsLimit(t *testing.T) {
	repo := &mocks.DailyUploadRepository{}
	pls := &mocks.PlanLogServiceUpload{}
	ps := &mocks.PlanServiceUpload{}
	us := &mocks.UserServiceUpload{}
	ls := &mocks.LogServiceUpload{}
	dailyUploadService := factories.NewDailyUploadService(repo, pls, ps, us, ls)
	repo.On("GetUserConsumption", mock.Anything).Return([]*psqlmodels.DailyUpload{
		{
			UploadSizeKb: 3000,
		},
	}, nil)
	pls.On("GetUserPlan", mock.Anything).Return(&psqlmodels.PlanLog{
		Plan: psqlmodels.Plan{
			UploadLimitSizeKb: 2000,
		},
	}, nil)

	// Call the function
	err := dailyUploadService.GetConsumption(123)

	// Check if the error message matches the expected error
	expectedErr := errors.New("total consumption limit has been breached")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}

func TestGetConsumption_DailyUploadLimitExceeded(t *testing.T) {
	repo := &mocks.DailyUploadRepository{}
	pls := &mocks.PlanLogServiceUpload{}
	ps := &mocks.PlanServiceUpload{}
	us := &mocks.UserServiceUpload{}
	ls := &mocks.LogServiceUpload{}
	dailyUploadService := factories.NewDailyUploadService(repo, pls, ps, us, ls)

	repo.On("GetUserConsumption", mock.Anything).Return([]*psqlmodels.DailyUpload{
		{
			ID:           1,
			UploadSizeKb: 1,
			CreatedAt:    time.Now(),
		},
		{
			ID:           2,
			UploadSizeKb: 1,
			CreatedAt:    time.Now(),
		},
		{
			ID:           3,
			UploadSizeKb: 1,
			CreatedAt:    time.Now(),
		},
	}, nil)

	pls.On("GetUserPlan", mock.Anything).Return(&psqlmodels.PlanLog{
		Plan: psqlmodels.Plan{
			DailyUploadLimit:  3,
			UploadLimitSizeKb: 2000,
		},
	}, nil)

	// Call the function
	err := dailyUploadService.GetConsumption(123)

	// Check if the error message matches the expected error
	expectedErr := errors.New("daily upload limit has been reached")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}

func TestGetConsumption_NoError(t *testing.T) {
	repo := &mocks.DailyUploadRepository{}
	pls := &mocks.PlanLogServiceUpload{}
	ps := &mocks.PlanServiceUpload{}
	us := &mocks.UserServiceUpload{}
	ls := &mocks.LogServiceUpload{}
	dailyUploadService := factories.NewDailyUploadService(repo, pls, ps, us, ls)

	repo.On("GetUserConsumption", mock.Anything).Return([]*psqlmodels.DailyUpload{
		{
			ID:           1,
			UploadSizeKb: 1,
			CreatedAt:    time.Now(),
		},
		{
			ID:           2,
			UploadSizeKb: 1,
			CreatedAt:    time.Now(),
		},
		{
			ID:           3,
			UploadSizeKb: 1,
		},
	}, nil)

	pls.On("GetUserPlan", mock.Anything).Return(&psqlmodels.PlanLog{
		Plan: psqlmodels.Plan{
			DailyUploadLimit:  3,
			UploadLimitSizeKb: 2000,
		},
	}, nil)

	err := dailyUploadService.GetConsumption(123)

	// Check if the returned error is nil
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}
