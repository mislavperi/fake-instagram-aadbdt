package repositories

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"github.com/mislavperi/fake-instagram-aadbdt/server/utils/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (r *UserRepository) Create(firstName string, lastName string, username string, email string, password string) error {
	user := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}
	result := r.Database.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) CheckCredentials(username string, password string) error {
	var result models.User
	if err := r.Database.First(&result).Where(&models.User{Username: username}).Error; err != nil {
		return err
	}
	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return errors.NewInvalidCredentialsError(err.Error())
	}
	return nil
}

func (r *UserRepository) FetchUserInformation(username string) (*models.User, error) {
	var result models.User
	if err := r.Database.First(&result).Where(&models.User{Username: username}).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
