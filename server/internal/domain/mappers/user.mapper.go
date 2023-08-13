package mappers

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) MapUserToDTO(plan models.Plan) psqlmodels.Plan {
	return psqlmodels.Plan{
		PlanName:          plan.PlanName,
		Cost:              plan.Cost,
		UploadLimitSizeKb: plan.UploadLimitSizeKb,
		DailyUploadLimit:  plan.DailyUploadLimit,
	}
}

func (m *UserMapper) MapGHUserToDTO(user models.GHUser, plan psqlmodels.Plan) psqlmodels.User {
	return psqlmodels.User{
		Email:    user.Email,
		Username: user.Username,
		Plan:     &plan,
	}
}

func (m *UserMapper) MapGoogleUserToDTO(user models.GoogleUser, plan psqlmodels.Plan) psqlmodels.User {
	return psqlmodels.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Email,
		Plan:      &plan,
	}
}

func (m *UserMapper) MapDTOToUser(user psqlmodels.User) models.User {
	return models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.Email,
		LastName:  user.LastName,
		Plan: models.Plan{
			PlanName:          user.Plan.PlanName,
			UploadLimitSizeKb: user.Plan.UploadLimitSizeKb,
			DailyUploadLimit:  user.Plan.DailyUploadLimit,
			Cost:              user.Plan.DailyUploadLimit,
		},
		Username: user.Username,
	}
}

func (m *UserMapper) MapUserToDTOO(user models.User) psqlmodels.User {
	return psqlmodels.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Role: &psqlmodels.Role{
			Name: user.Role.Name,
		},
		Plan: &psqlmodels.Plan{
			PlanName:          user.Plan.PlanName,
			UploadLimitSizeKb: user.Plan.UploadLimitSizeKb,
			DailyUploadLimit:  user.Plan.DailyUploadLimit,
			Cost:              user.Plan.Cost,
		},
	}
}
