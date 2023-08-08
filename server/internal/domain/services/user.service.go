package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"

	"golang.org/x/crypto/bcrypt"
)

type UserMapper interface {
	MapUserToDTO(plan models.Plan) psqlmodels.Plan
}

type UserRepository interface {
	Create(firstName string, lastName string, username string, email string, password string) error
	CheckCredentials(username string, password string) error
	FetchUserInformation(username string) (*psqlmodels.User, error)
	SetUserPlan(username string, plan psqlmodels.Plan) error
}

type UserService struct {
	UserRepository UserRepository
	userMapper     UserMapper
}

const SECRET_KEY = "f83edb0a3b4e9547fd6fbd981513bce0d604472c547daaeed8907a78c5793671"

func NewUserService(userRepository UserRepository, userMapper UserMapper) *UserService {
	return &UserService{
		UserRepository: userRepository,
		userMapper:     userMapper,
	}
}

func (s *UserService) GetUserInformation(username string) (*models.User, error) {
	user, err := s.UserRepository.FetchUserInformation(username)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (s *UserService) Create(firstName string, lastName string, username string, email string, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	err = s.UserRepository.Create(firstName, lastName, username, email, string(bytes))
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(username string, password string) (*string, *string, error) {
	err := s.UserRepository.CheckCredentials(username, password)
	if err != nil {
		return nil, nil, err
	}

	accessToken, refreshToken, err := s.generateTokenPair(username)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) generateTokenPair(username string) (*string, *string, error) {
	accessExpirationTime := time.Now().Add(5 * time.Minute).Unix()
	refreshExpirationTime := time.Now().Add(45 * time.Minute).Unix()

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Identifier: username,
		Type:       "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime,
		},
	}).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Identifier: username,
		Type:       "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime,
		},
	}).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, nil, nil
	}

	return &accessToken, &refreshToken, nil
}

func (s *UserService) SelectUserPlan(username string, plan models.Plan) error {
	mappedPlan := s.userMapper.MapUserToDTO(plan)
	err := s.UserRepository.SetUserPlan(username, mappedPlan)
	if err != nil {
		return nil
	}
	return nil
}
