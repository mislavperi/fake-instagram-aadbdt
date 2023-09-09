package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
)

//go:generate mockery --output=./tests/mocks --name=UploadService
type UploadService interface {
	GetConsumption(userID int) error
	GetStatistics(userID int) (*models.Plan, *uint64, *int, *int, error)
	GetExpandedStatistics(userID int) (*models.User, *models.Plan, *uint64, *int, *int, error)
}

type UploadController struct {
	uploadService UploadService
}

func NewUploadController(uploadService UploadService) *UploadController {
	return &UploadController{
		uploadService: uploadService,
	}
}

func (c *UploadController) GetStatistics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		identifier, err := strconv.Atoi(ctx.GetHeader("identifier"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		plan, totalConsumptionKb, totalDailyUploadCount, totalConsumptionCount, err := c.uploadService.GetStatistics(identifier)
		if err != nil {

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}

		statisticResponse := models.Statistics{
			Plan:                  *plan,
			TotalConsumptionKb:    *totalConsumptionKb,
			TotalDailyUploadCount: *totalDailyUploadCount,
			TotalConsumptionCount: *totalConsumptionCount,
		}

		ctx.JSON(http.StatusOK, statisticResponse)
	}
}

func (c *UploadController) GetUserStatistics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userID int
		requestId := ctx.Query("id")
		err := json.Unmarshal([]byte(requestId), &userID)
		if err != nil {

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		user, plan, totalConsumptionKb, totalDailyUploadCount, totalConsumptionCount, err := c.uploadService.GetExpandedStatistics(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}

		statisticResponse := models.ExpandedStatistics{
			User:                  *user,
			Plan:                  *plan,
			TotalConsumptionKb:    *totalConsumptionKb,
			TotalDailyUploadCount: *totalDailyUploadCount,
			TotalConsumptionCount: *totalConsumptionCount,
		}

		ctx.JSON(http.StatusOK, statisticResponse)
	}
}
