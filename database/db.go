package database

import (
	"fmt"
	"os"

	"github.com/Tooluloope/gourl/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return &Database{}, fmt.Errorf("error opening db connection: %w", err)
	}

	fmt.Println("Database Connected!")

	return &Database{
		Client: db,
	}, nil

}

func Migrate(database *Database) (*Database, error) {
	err := database.Client.AutoMigrate(&models.User{}, &models.URL{})

	if err != nil {
		return database, fmt.Errorf("error migrating database: %w", err)
	}
	return database, nil
}
