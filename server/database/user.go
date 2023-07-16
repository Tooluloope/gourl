package database

import (
	"context"
	"errors"

	"github.com/Tooluloope/gourl/server/models"
	"gorm.io/gorm"
)

func (db *Database) AuthenticateUser(ctx context.Context, email string) (models.User, error) {

	user := models.User{
		Email: email,
	}

	if result := db.Client.Where("email = ?", email).First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("Invalid email or password")
		}
		return models.User{}, result.Error
	}
	return user, nil
}

func (db *Database) RegisterUser(ctx context.Context, user models.User) (models.User, error) {

	result := db.Client.Where("email = ?", user.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

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
