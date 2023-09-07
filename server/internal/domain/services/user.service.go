package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	enums "github.com/mislavperi/fake-instagram-aadbdt/server/utils/enums/action"
	"github.com/prometheus/client_golang/prometheus"

	"golang.org/x/crypto/bcrypt"
)

type UserMetrics interface {
	OnLoginStart(label string) *prometheus.Timer
	OnLoginFinish(timer *prometheus.Timer)
	OnCreationStart(label string) *prometheus.Timer
	OnCreationFinish(timer *prometheus.Timer)
}

type UserMapper interface {
	MapUserToDTO(plan models.Plan) psqlmodels.Plan
	MapGHUserToDTO(user models.GHUser) psqlmodels.User
	MapGoogleUserToDTO(user models.GoogleUser) psqlmodels.User
	MapDTOToUser(user psqlmodels.User) models.User
	MapUserToDTOO(user models.User) psqlmodels.User
	MapDTOToUsers(users []*psqlmodels.User) []models.User
}

type UserRepository interface {
	Create(firstName string, lastName string, username string, email string, password string) (*int, error)
	CheckCredentials(username string, password string) (*int, error)
	FetchUserInformation(id int) (*psqlmodels.User, error)
	AuthenticateGithubUser(psqlmodels.User) (*int, error)
	AuthenticateGoogleUser(psqlmodels.User) (*int, error)
	GetAllUsers() ([]*psqlmodels.User, error)
}

type PlanDomain interface {
	GetPlan(planName string) (*psqlmodels.Plan, error)
}

type UserService struct {
	planLogService *PlanLogService
	logService     *LogService

	userMapper UserMapper

	UserRepository UserRepository

	metrics UserMetrics

	ghClientId     string
	ghClientSecret string
	secretKey      string
}

func NewUserService(userRepository UserRepository, userMapper UserMapper, planService PlanDomain, planLogService *PlanLogService, logService *LogService, metrics UserMetrics, ghClientId string, ghClientSecret string, secretKey string) *UserService {
	return &UserService{
		UserRepository: userRepository,
		userMapper:     userMapper,
		planLogService: planLogService,
		logService:     logService,
		ghClientId:     ghClientId,
		ghClientSecret: ghClientSecret,
		secretKey:      secretKey,
		metrics:        metrics,
	}
}

func (s *UserService) MapUserToDTO(user models.User) psqlmodels.User {
	return s.userMapper.MapUserToDTOO(user)
}

func (s *UserService) GetUserInformation(id int) (*models.User, error) {
	user, err := s.UserRepository.FetchUserInformation(id)
	if err != nil {
		return nil, err
	}
	mappedUser := s.userMapper.MapDTOToUser(*user)
	s.logService.LogAction(id, enums.GET_USER_INFO.String())
	return &mappedUser, nil
}

func (s *UserService) Create(firstName string, lastName string, username string, email string, password string) error {
	timer := s.metrics.OnCreationStart("standard_user")
	defer s.metrics.OnCreationFinish(timer)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	userID, err := s.UserRepository.Create(firstName, lastName, username, email, string(bytes))
	if err != nil {
		return err
	}
	s.logService.LogAction(*userID, enums.CREATE_USER.String())
	return nil
}

func (s *UserService) Login(username string, password string) (*string, *string, error) {
	timer := s.metrics.OnLoginStart("standard_user")
	defer s.metrics.OnLoginFinish(timer)
	id, err := s.UserRepository.CheckCredentials(username, password)
	if err != nil {
		return nil, nil, err
	}

	accessToken, refreshToken, err := s.generateTokenPair(*id)
	if err != nil {
		return nil, nil, err
	}

	s.logService.LogAction(*id, enums.LOGIN_USER.String())
	return accessToken, refreshToken, nil
}

func (s *UserService) generateTokenPair(userID int) (*string, *string, error) {
	accessExpirationTime := time.Now().Add(5 * time.Minute).Unix()
	refreshExpirationTime := time.Now().Add(45 * time.Minute).Unix()

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Identifier: strconv.Itoa(userID),
		Type:       "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime,
		},
	}).SignedString([]byte(s.secretKey))
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Identifier: strconv.Itoa(userID),
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

func (s *UserService) SelectUserPlan(id int, plan models.Plan) error {
	err := s.planLogService.InsertPlanChangeLog(int64(id), plan.ID)
	if err != nil {
		return err
	}
	s.logService.LogAction(id, enums.CHANGE_PLAN.String())
	return nil
}

func (s *UserService) InsertAdminPlanChange(adminID int, userID int, planID int) error {
	err := s.planLogService.InsertAdminPlanChangeLog(int64(userID), int64(planID))
	if err != nil {
		return err
	}
	s.logService.LogAction(adminID, enums.CHANGE_PLAN.String())
	return nil
}

func (s *UserService) AuthenticateGithubUser(credentials models.GHCredentials) (*string, *string, error) {
	var ghToken models.GHToken
	var ghUser models.GHUser
	timer := s.metrics.OnLoginStart("github_user")
	defer s.metrics.OnLoginFinish(timer)

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
	mappedUser := s.userMapper.MapGHUserToDTO(ghUser)
	if err != nil {
		return nil, nil, err
	}
	id, err := s.UserRepository.AuthenticateGithubUser(mappedUser)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	accessToken, refreshToken, err := s.generateTokenPair(*id)
	if err != nil {
		return nil, nil, err
	}
	s.logService.LogAction(*id, enums.LOGIN_USER.String())
	return accessToken, refreshToken, nil
}

func (s *UserService) AuthenticateGoogleUser(credentials models.GoogleToken) (*string, *string, error) {
	var googleUser models.GoogleUser
	timer := s.metrics.OnLoginStart("google_user")
	defer s.metrics.OnLoginFinish(timer)

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
	mappedUser := s.userMapper.MapGoogleUserToDTO(*claims)
	id, err := s.UserRepository.AuthenticateGoogleUser(mappedUser)
	if err != nil {
		return nil, nil, err
	}
	accessToken, refreshToken, err := s.generateTokenPair(mappedUser.ID)
	if err != nil {
		return nil, nil, err
	}
	s.logService.LogAction(*id, enums.LOGIN_USER.String())
	return accessToken, refreshToken, nil
}

func (s *UserService) GetAllUsers(adminID int) ([]models.User, error) {
	users, err := s.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	mappedUsers := s.userMapper.MapDTOToUsers(users)
	s.logService.LogAction(adminID, enums.GET_USERS.String())
	return mappedUsers, nil
}

func (s *UserService) GetUserLogs(userID int, adminID int) ([]models.Log, error) {
	logs, err := s.logService.GetUserLogs(userID)
	if err != nil {
		return nil, err
	}
	s.logService.LogAction(adminID, enums.GET_USER_LOGS.String())
	return logs, nil
}
