package bootstrap

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/cmd/api/config"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/controllers"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/mappers"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/metrics"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/repositories"
	repository "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/s3"
)

func Api() (*api.API, error) {
	db, err := psql.NewDatabaseConnection(config.Cfg.Database.Host, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Name, config.Cfg.Database.Port)
	logRepository := repositories.NewLogRepository(db)
	logMapper := mappers.NewLogMapper()
	logService := services.NewLogService(logRepository, logMapper)
	planLogRepository := repositories.NewPlanLogRepository(db)
	planLogService := services.NewPlanLogService(planLogRepository)

	userMetrics := metrics.NewUserMetrics("fake-instagram")

	planRepository := repositories.NewPlanRepository(db)
	planMapper := mappers.NewPlanMapper()
	planService := services.NewPlanService(planRepository, planMapper, logRepository)
	userMapper := mappers.NewUserMapper()
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, userMapper, planService, planLogService, logService, userMetrics, config.Cfg.Github.ClientID, config.Cfg.Github.ClientSecret, config.Cfg.Auth.SecretKey)
	uploadRepository := repositories.NewDailyUploadRepository(db)
	uploadService := services.NewDailyUploadService(uploadRepository, planLogService, planService, userService)
	s3Repository := repository.NewS3Repository(config.Cfg.Aws.Bucket, config.Cfg.Aws.Region, config.Cfg.Aws.AccessKeyId, config.Cfg.Aws.SecretAccessKey, "")

	planController := controllers.NewPlanController(planService)

	pictureMetrics := metrics.NewPictureMetrics("fake-instagram")

	userController := controllers.NewUserController(userService)
	pictureMapper := mappers.NewPictureMapper()
	pictureRepository := repositories.NewPictureRepository(db)
	pictureService := services.NewPictureService(userService, uploadService, pictureRepository, pictureMapper, logRepository, s3Repository, pictureMetrics)
	pictureController := controllers.NewPictureController(pictureService)
	uploadController := controllers.NewUploadController(uploadService)

	if err != nil {
		panic(err)
	}
	api := api.NewAPI(userController, planController, pictureController, uploadController, 8080)
	return api, nil
}
