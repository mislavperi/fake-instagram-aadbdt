package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/controllers"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/controllers/tests/mocks"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

func TestRoutes(t *testing.T) {
	userService := &mocks.UserService{}
	planService := &mocks.PlanService{}
	pictureService := &mocks.PictureService{}
	uploadService := &mocks.UploadService{}

	a := api.NewAPI(
		controllers.NewUserController(userService),
		controllers.NewPlanController(planService),
		controllers.NewPictureController(pictureService),
		controllers.NewUploadController(uploadService),
		8080,
	)

	t.Run("TestPlanRoute", func(t *testing.T) {
		planService.On("GetPlans").Return([]models.Plan{}, nil)
		req, err := http.NewRequest("GET", "/api/plans/get", nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		a.Gin.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("expected status 200 but got %d", recorder.Code)
		}
	})

	t.Run("TestPictures", func(t *testing.T) {
		pictureService.On("GetImages", mock.Anything).Return([]models.Picture{}, nil)
		req, err := http.NewRequest("GET", "/api/public/picture/get?filter={}", nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		a.Gin.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("expected status 200 but got %d", recorder.Code)
		}
	})
}
