package psql

import (
	"fmt"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(
	host string,
	user string,
	password string,
	databaseName string,
	port string,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, databaseName, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Package{})
	return db, nil
}
