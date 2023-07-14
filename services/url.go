package services

import (
	"context"

	"github.com/Tooluloope/gourl/models"
)

func (service *Service) CreateURL(ctx context.Context, url models.URL) (models.URL, error) {

	url, err := service.store.CreateURL(ctx, url)

	if err != nil {
		return models.URL{}, err
	}

	return url, nil
}

func (service *Service) GetURLByShortCode(ctx context.Context, shortCode string) (models.URL, error) {
	return models.URL{}, nil
}

func (service *Service) GetAllURLs(ctx context.Context) ([]models.URL, error) {
	return nil, nil
}

func (service *Service) DeleteURL(ctx context.Context, shortCode string) error {
	return nil
}

func (service *Service) UpdateURL(ctx context.Context, url models.URL) (models.URL, error) {
	return models.URL{}, nil
}
