package infrastructure

import (
	"errors"
	"fmt"
	"log"

	"github.com/ronnachate/inventory-api-go/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database", err)
	}

	MigrateDB()
}

func CloseDBConnection() {
	dbInstance, _ := DB.DB()
	_ = dbInstance.Close()
}

// Need to be refractor later
// https://gorm.io/docs/migration.html
func MigrateDB() {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(&domain.User{}, &domain.UserStatus{})
	statuses := []domain.UserStatus{
		{Name: "Active", ID: 1},
		{Name: "Inactive", ID: 2},
		{Name: "Deleted", ID: 3},
	}
	if err := DB.First(&domain.UserStatus{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		DB.Create(statuses)
	}
}
