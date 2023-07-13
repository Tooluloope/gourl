package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/Tooluloope/gourl/models"
	"gorm.io/gorm"
)

func (db *Database) AuthenticateUser(ctx context.Context, username, password string) (models.User, error) {
	return models.User{}, nil
}

func (db *Database) RegisterUser(ctx context.Context, user models.User) (models.User, error) {

	result := db.Client.Where("email = ?", user.Email).First(&user)
	fmt.Println(result)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		fmt.Println("User not found, creating new user")
		if err := user.HashPassword(); err != nil {
			return models.User{}, err
		}

		if err := db.Client.Create(&user).Error; err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	return models.User{}, gorm.ErrDuplicatedKey

}
