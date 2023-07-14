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

	result := db.Client.Where("short_url = ?", url.ShortURL).First(&url)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		if err := db.Client.Create(&url).Error; err != nil {
			return models.URL{}, err
		}
		return url, nil
	}

	return models.URL{}, gorm.ErrDuplicatedKey
}

func (db *Database) GetURLByShortCode(ctx context.Context, shortCode string) (models.URL, error) {
	return models.URL{}, nil
}

func (db *Database) GetAllURLs(ctx context.Context) ([]models.URL, error) {
	return nil, nil
}

func (db *Database) DeleteURL(ctx context.Context, shortCode string) error {
	return nil
}

func (db *Database) UpdateURL(ctx context.Context, url models.URL, shortCode string) error {
	return nil
}
