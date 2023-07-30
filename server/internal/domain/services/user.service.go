package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(firstName string, lastName string, username string, email string, password string) error
	CheckCredentials(username string, password string) error
}

type UserService struct {
	UserRepository UserRepository
}

const SECRET_KEY = "akj fas"

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) Create(firstName string, lastName string, username string, email string, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
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
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Identifier:     username,
		Type:           "access",
		StandardClaims: jwt.StandardClaims{},
	}).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Identifier:     username,
		Type:           "refresh",
		StandardClaims: jwt.StandardClaims{},
	}).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, nil, nil
	}

	return &accessToken, &refreshToken, nil
}
