package database

import (
	"context"

	"github.com/Tooluloope/gourl/models"
)

func (db *Database) CreateURL(ctx context.Context, url models.URL, shortCode string) error {
	return nil
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
