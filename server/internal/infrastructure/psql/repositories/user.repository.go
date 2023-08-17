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

func (r *UserRepository) Create(firstName string, lastName string, username string, email string, password string) (*int, error) {
	user := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}
	if err := r.Database.Create(&user).Error; err != nil {
		return nil, err
	}

	var userObj models.User

	if err := r.Database.Last(&userObj).Error; err != nil {
		return nil, err
	}
	return &userObj.ID, nil
}

func (r *UserRepository) CheckCredentials(username string, password string) (*int, error) {
	var result models.User
	if err := r.Database.First(&result).Where(&models.User{Username: username}).Error; err != nil {
		return nil, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return nil, customerrors.NewInvalidCredentialsError(err.Error())
	}
	return &result.ID, err
}

func (r *UserRepository) FetchUserInformation(id int) (*models.User, error) {
	var result models.User
	if err := r.Database.Preload("Role").First(&result).Where("id = ?", id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepository) AuthenticateGithubUser(mappedUser models.User) (*int, error) {
	var result *models.User
	if err := r.Database.First(&result).Where(&models.User{Username: mappedUser.Username}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if err = r.Database.Create(&mappedUser).Error; err != nil {
			return nil, err
		}
		var latestUser models.User
		if err := r.Database.Last(&latestUser).Error; err != nil {
			if err = r.Database.Create(&mappedUser).Error; err != nil {
				return nil, err
			}
		}
		return &latestUser.ID, nil
	}
	return &result.ID, nil
}

func (r *UserRepository) AuthenticateGoogleUser(mappedUser models.User) (*int, error) {
	var result *models.User
	if err := r.Database.First(&result).Where(&models.User{Username: mappedUser.Username}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if res := r.Database.Create(&mappedUser); res.Error != nil {
			return nil, err
		}
		var latestUser models.User
		if err := r.Database.Last(&latestUser).Error; err != nil {
			if err = r.Database.Create(&mappedUser).Error; err != nil {
				return nil, err
			}
		}
		return &latestUser.ID, nil
	}
	return &result.ID, nil
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := r.Database.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
