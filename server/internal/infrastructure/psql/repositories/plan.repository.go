package repositories

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"gorm.io/gorm"
)

type PlanRepository struct {
	Database *gorm.DB
}

func NewPlanRepository(database *gorm.DB) *PlanRepository {
	return &PlanRepository{
		Database: database,
	}
}

func (r *PlanRepository) GetPlans() ([]models.Plan, error) {
	var plans []models.Plan
	if err := r.Database.Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}