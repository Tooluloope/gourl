package services

import (
	"context"

	"github.com/Tooluloope/gourl/server/models"
)

type Store interface {
	AuthenticateUser(ctx context.Context, email string) (models.User, error)
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
	GetURLByShortCode(ctx context.Context, shortCode string) (models.URL, error)
	CreateURL(ctx context.Context, url models.URL) (models.URL, error)
	GetAllURLs(ctx context.Context) ([]models.URL, error)
	DeleteURL(ctx context.Context, urlId string) error
	UpdateURL(ctx context.Context, url models.URL) (models.URL, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}
