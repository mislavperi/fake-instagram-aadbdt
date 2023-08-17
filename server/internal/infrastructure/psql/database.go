package psql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB
var err error

func NewDatabaseConnection(
	host string,
	user string,
	password string,
	databaseName string,
	port string,
) (*gorm.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, databaseName, port,
	)
	dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	dbInstance.AutoMigrate(&models.User{}, &models.Role{}, &models.Plan{}, &models.Picture{}, &models.Log{}, &models.DailyUpload{}, &models.PlanLog{})
	return dbInstance, nil
}
