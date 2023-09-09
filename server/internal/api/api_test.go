package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
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
		plan := models.Plan{
			ID:                1,
			PlanName:          "PlanOne",
			UploadLimitSizeKb: 1,
			DailyUploadLimit:  1,
			Cost:              1,
		}

		planService.On("GetPlans").Return([]models.Plan{
			plan,
		}, nil)
		req, err := http.NewRequest("GET", "/api/plans/get", nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		a.Gin.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("expected status 200 but got %d", recorder.Code)
		}
		var expectedPlan []models.Plan
		json.NewDecoder(recorder.Body).Decode(&expectedPlan)

		if !reflect.DeepEqual(expectedPlan, []models.Plan{plan}) {
			t.Errorf("expected plan to be same, but it isn't")
		}
	})

	t.Run("TestPictures", func(t *testing.T) {
		picture := models.Picture{
			Title: "Image",
		}

		pictureService.On("GetImages", mock.Anything).Return([]models.Picture{picture}, nil)
		req, err := http.NewRequest("GET", "/api/public/picture/get?filter={}", nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		a.Gin.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("expected status 200 but got %d", recorder.Code)
		}
		var expectedImage []models.Picture
		json.NewDecoder(recorder.Body).Decode(&expectedImage)
		if !reflect.DeepEqual(expectedImage, []models.Picture{picture}) {
			t.Errorf("expected plan to be same, but it isn't")
		}
	})
}
