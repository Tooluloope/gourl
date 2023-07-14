package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/Tooluloope/gourl/models"
	"github.com/Tooluloope/gourl/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (db *Database) CreateURL(ctx context.Context, url models.URL) (models.URL, error) {

	user := models.User{}
	id, ok := ctx.Value(utils.ContextKeyUser).(string)

	if !ok {
		fmt.Println(id)
		return models.URL{}, errors.New("Error creating url")
	}

	userId, err := uuid.Parse(id)

	if err != nil {
		return models.URL{}, errors.New("Error creating url")
	}

	if result := db.Client.Where("id = ?", userId).First(&user); result.Error != nil {
		fmt.Println(result.Error)
		return models.URL{}, errors.New("Error creating url")
	}

	url.User = user
	url.UserID = user.ID

	result := db.Client.Where("short_code = ?", url.ShortCode).First(&url)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		if err := db.Client.Create(&url).Error; err != nil {
			return models.URL{}, err
		}
		return url, nil
	}

	return models.URL{}, gorm.ErrDuplicatedKey
}

func (db *Database) GetURLByShortCode(ctx context.Context, shortCode string) (models.URL, error) {

	url := models.URL{}

	if result := db.Client.Where("short_code = ?", shortCode).First(&url); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.URL{}, errors.New("Invalid short code")
		}
		return models.URL{}, result.Error
	}

	return url, nil
}

func (db *Database) GetAllURLs(ctx context.Context) ([]models.URL, error) {

	var user models.User

	if result := db.Client.Preload("URLs").First(&user, "id = ?", ctx.Value(utils.ContextKeyUser)); result.Error != nil {
		return []models.URL{}, result.Error
	}

	return user.URLs, nil
}

func (db *Database) DeleteURL(ctx context.Context, urlId string) error {

	var url models.URL
	userId := ctx.Value(utils.ContextKeyUser)

	if result := db.Client.Where("id = ? AND user_id = ?", urlId, userId).First(&url); result.Error != nil {
		return result.Error
	}

	if userId != url.UserID {
		if result := db.Client.Unscoped().Delete(&url); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (db *Database) UpdateURL(ctx context.Context, url models.URL) (models.URL, error) {

	var urlToUpdate models.URL
	userId := ctx.Value(utils.ContextKeyUser)

	if result := db.Client.Where("id = ? AND user_id = ?", url.ID, userId).First(&urlToUpdate); result.Error != nil {
		return models.URL{}, result.Error
	}

	if userId != urlToUpdate.UserID {
		if result := db.Client.Model(&urlToUpdate).Updates(&url); result.Error != nil {
			return models.URL{}, result.Error
		}
	}

	return urlToUpdate, nil
}
