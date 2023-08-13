package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services/interfaces"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	enums "github.com/mislavperi/fake-instagram-aadbdt/server/utils/enums/action"

	"golang.org/x/crypto/bcrypt"
)

type UserMapper interface {
	MapUserToDTO(plan models.Plan) psqlmodels.Plan
	MapGHUserToDTO(user models.GHUser, plan psqlmodels.Plan) psqlmodels.User
	MapGoogleUserToDTO(user models.GoogleUser, plan psqlmodels.Plan) psqlmodels.User
	MapDTOToUser(user psqlmodels.User) models.User
	MapUserToDTOO(user models.User) psqlmodels.User
}

type UserRepository interface {
	Create(firstName string, lastName string, username string, email string, password string) error
	CheckCredentials(username string, password string) error
	FetchUserInformation(username string) (*psqlmodels.User, error)
	SetUserPlan(username string, plan psqlmodels.Plan) error
	AuthenticateGithubUser(psqlmodels.User) error
	AuthenticateGoogleUser(psqlmodels.User) error
}

type PlanDomain interface {
	GetPlan(planName string) (*psqlmodels.Plan, error)
}

type UserService struct {
	planService PlanDomain

	userMapper UserMapper

	UserRepository UserRepository
	logRepository  interfaces.LogRepository

	ghClientId     string
	ghClientSecret string
	secretKey      string
}

func NewUserService(userRepository UserRepository, userMapper UserMapper, planService PlanDomain, logRepository interfaces.LogRepository, ghClientId string, ghClientSecret string, secretKey string) *UserService {
	return &UserService{
		UserRepository: userRepository,
		userMapper:     userMapper,
		planService:    planService,
		logRepository:  logRepository,
		ghClientId:     ghClientId,
		ghClientSecret: ghClientSecret,
		secretKey:      secretKey,
	}
}

func (s *UserService) MapUserToDTO(user models.User) psqlmodels.User {
	return s.userMapper.MapUserToDTOO(user)
}

func (s *UserService) GetUserInformation(username string) (*models.User, error) {
	user, err := s.UserRepository.FetchUserInformation(username)
	if err != nil {
		return nil, err
	}
	mappedUser := s.userMapper.MapDTOToUser(*user)
	return &mappedUser, nil
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
	s.logRepository.LogAction(nil, enums.CREATE_USER.String())
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

	user, err := s.UserRepository.FetchUserInformation(username)
	if err != nil {
		return nil, nil, err
	}

	s.logRepository.LogAction(user, enums.LOGIN_USER.String())
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
	}).SignedString([]byte(s.secretKey))
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Identifier: username,
		Type:       "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime,
		},
	}).SignedString([]byte(s.secretKey))
	if err != nil {
		return nil, nil, nil
	}

	return &accessToken, &refreshToken, nil
}

func (s *UserService) SelectUserPlan(username string, plan models.Plan) error {
	mappedPlan := s.userMapper.MapUserToDTO(plan)
	err := s.UserRepository.SetUserPlan(username, mappedPlan)
	if err != nil {
		return err
	}

	user, err := s.UserRepository.FetchUserInformation(username)
	if err != nil {
		return err
	}

	s.logRepository.LogAction(user, enums.LOGIN_USER.String())
	return nil
}

func (s *UserService) AuthenticateGithubUser(credentials models.GHCredentials) (*string, *string, error) {
	var ghToken models.GHToken
	var ghUser models.GHUser

	body := models.GHCredsReq{
		Code:         credentials.Code,
		ClientID:     s.ghClientId,
		ClientSecret: s.ghClientSecret,
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(bodyJSON),
	)
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&ghToken)
	if err != nil {
		return nil, nil, err
	}
	infoRequest, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, nil, err
	}
	infoRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ghToken.AccessToken))
	infoResp, err := http.DefaultClient.Do(infoRequest)
	if err != nil {
		return nil, nil, err
	}
	err = json.NewDecoder(infoResp.Body).Decode(&ghUser)
	if err != nil || ghUser.Username == "" {
		return nil, nil, err
	}
	mappedPlan, err := s.planService.GetPlan("FREE")

	mappedUser := s.userMapper.MapGHUserToDTO(ghUser, *mappedPlan)
	if err != nil {
		return nil, nil, err
	}
	err = s.UserRepository.AuthenticateGithubUser(mappedUser)
	if err != nil {
		return nil, nil, err
	}
	accessToken, refreshToken, err := s.generateTokenPair(mappedUser.Username)
	if err != nil {
		return nil, nil, err
	}
	s.logRepository.LogAction(&mappedUser, enums.LOGIN_USER.String())
	return accessToken, refreshToken, nil
}

func (s *UserService) AuthenticateGoogleUser(credentials models.GoogleToken) (*string, *string, error) {
	var googleUser models.GoogleUser

	token, err := jwt.ParseWithClaims(
		credentials.GoogleJWT,
		&googleUser,
		func(t *jwt.Token) (interface{}, error) {
			resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
			if err != nil {
				return nil, err
			}
			dat, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			cemKey := map[string]string{}
			err = json.Unmarshal(dat, &cemKey)
			if err != nil {
				return nil, err
			}
			pem, ok := cemKey[fmt.Sprintf("%s", t.Header["kid"])]
			if !ok {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(*models.GoogleUser)
	if !ok {
		return nil, nil, err
	}
	mappedPlan, err := s.planService.GetPlan("FREE")
	if err != nil {
		return nil, nil, err
	}
	mappedUser := s.userMapper.MapGoogleUserToDTO(*claims, *mappedPlan)
	err = s.UserRepository.AuthenticateGoogleUser(mappedUser)
	if err != nil {
		return nil, nil, err
	}
	accessToken, refreshToken, err := s.generateTokenPair(mappedUser.Email)
	if err != nil {
		return nil, nil, err
	}
	s.logRepository.LogAction(&mappedUser, enums.LOGIN_USER.String())
	return accessToken, refreshToken, nil
}
