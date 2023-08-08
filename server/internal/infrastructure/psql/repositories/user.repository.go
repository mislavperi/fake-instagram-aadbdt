package repositories

import (
	"errors"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	customerrors "github.com/mislavperi/fake-instagram-aadbdt/server/utils/errors"
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
		return customerrors.NewInvalidCredentialsError(err.Error())
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

func (r *UserRepository) SetUserPlan(username string, plan models.Plan) error {
	var result models.User
	if err := r.Database.First(&result).Where(&models.User{Username: username}).Error; err != nil {
		return err
	}
	result.Plan = plan
	r.Database.Save(&result)
	return nil
}

func (r *UserRepository) AuthenticateGithubUser(mappedUser models.User) error {
	var result *models.User
	if err := r.Database.First(&result).Where(&models.User{Username: mappedUser.Username}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if res := r.Database.Create(&mappedUser); res.Error != nil {
			return err
		}
		return nil
	}
	return nil
}

func (r *UserRepository) AuthenticateGoogleUser(mappedUser models.User) error {
	var result *models.User
	if err := r.Database.First(&result).Where(&models.User{Username: mappedUser.Username}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if res := r.Database.Create(&mappedUser); res.Error != nil {
			return err
		}
		return nil
	}
	return nil
}
