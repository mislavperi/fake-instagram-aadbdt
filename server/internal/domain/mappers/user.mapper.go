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

func (m *UserMapper) MapGHUserToDTO(user models.GHUser) psqlmodels.User {
	return psqlmodels.User{
		Email:    user.Email,
		Username: user.Username,
	}
}

func (m *UserMapper) MapGoogleUserToDTO(user models.GoogleUser) psqlmodels.User {
	return psqlmodels.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Email,
	}
}

func (m *UserMapper) MapDTOToUser(user psqlmodels.User) models.User {
	return models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.Email,
		LastName:  user.LastName,
		Username:  user.Username,
		RoleID:    user.RoleID,
	}
}

func (m *UserMapper) MapDTOToUsers(users []*psqlmodels.User) []models.User {
	var mappedUsers []models.User
	for _, user := range users {
		mappedUsers = append(mappedUsers, models.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			RoleID:    user.RoleID,
			Username:  user.Username,
			Email:     user.Email,
		})
	}

	return mappedUsers
}

func (m *UserMapper) MapUserToDTOO(user models.User) psqlmodels.User {
	return psqlmodels.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		RoleID:    user.RoleID,
	}
}
