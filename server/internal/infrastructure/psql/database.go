package psql

import (
	"fmt"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, databaseName, port,
	)
	dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	dbInstance.AutoMigrate(&models.User{}, &models.Role{}, &models.Plan{})
	return dbInstance, nil
}
