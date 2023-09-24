package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
)

type PlanService interface {
	GetPlans() ([]models.Plan, error)
}

//go:generate mockery --output=./tests/mocks --name=PlanService
type PlanController struct {
	planService PlanService
}

func NewPlanController(planService PlanService) *PlanController {
	return &PlanController{
		planService: planService,
	}
}

func (c *PlanController) GetPlans() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		plans, err := c.planService.GetPlans()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, plans)
	}
}
