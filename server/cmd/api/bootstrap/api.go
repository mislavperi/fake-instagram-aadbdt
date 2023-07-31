package bootstrap

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/cmd/api/config"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/api/controllers"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/mappers"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/repositories"
)

func Api() (*api.API, error) {
	db, err := psql.NewDatabaseConnection(config.Cfg.Database.Host, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Name, config.Cfg.Database.Port)
	
	userMapper := mappers.NewUserMapper()
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, userMapper)
	userController := controllers.NewUserController(userService)

	planMapper := mappers.NewPlanMapper()
	planRepository := repositories.NewPlanRepository(db)
	planService := services.NewPlanService(planRepository, planMapper)
	planController := controllers.NewPlanController(planService)

	if err != nil {
		panic(err)
	}
	api := api.NewAPI(userController, planController, 8080)
	return api, nil
}
