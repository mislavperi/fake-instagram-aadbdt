package repositories

import (
	"errors"
	"time"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"gorm.io/gorm"
)

type PlanLogRepository struct {
	database *gorm.DB
}

const (
	FREE_PLAN_ID = 1
)

func NewPlanLogRepository(database *gorm.DB) *PlanLogRepository {
	return &PlanLogRepository{
		database: database,
	}
}

func (r *PlanLogRepository) InsertPlanChange(userID int64, planID int64) error {
	var planLogs []models.PlanLog
	today := time.Now()

	err := r.database.Where("user_id = ?", userID).Find(&planLogs).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	for _, planLog := range planLogs {
		if planLog.CreatedAt.Year() == today.Year() && today.Month() == planLog.CreatedAt.Month() && planLog.CreatedAt.Day() == today.Day() {
			return errors.New("you've already changed the plan today")
		}
	}
	if err := r.database.Create(&models.PlanLog{PlanID: planID, UserID: userID}).Error; err != nil {
		return err
	}
	return nil
}

func (r *PlanLogRepository) InsertAdminPlanChange(userID int64, planID int64) error {
	if err := r.database.Create(&models.PlanLog{PlanID: planID, UserID: userID}).Error; err != nil {
		return err
	}
	return nil
}

func (r *PlanLogRepository) GetUserPlan(userID int64) (*models.PlanLog, error) {
	var userPlanLog models.PlanLog
	if err := r.database.Preload("Plan").Where("user_id = ?", userID).Last(&userPlanLog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.database.Create(&models.PlanLog{
				UserID:    userID,
				PlanID:    1,
				CreatedAt: time.Now(),
			})
			var latestUserPlanEntry models.PlanLog
			r.database.Preload("Plan").Where("user_id = ?", userID).Last(&latestUserPlanEntry)
			return &latestUserPlanEntry, nil
		} else {
			return nil, err
		}
	}
	return &userPlanLog, nil
}
