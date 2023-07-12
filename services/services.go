package services

import (
	"context"

	"github.com/Tooluloope/gourl/models"
)

type Store interface {
	AuthenticateUser(ctx context.Context, username, password string) (models.User, error)
	RegisterUser(ctx context.Context, user models.User) error
	GetURLByShortCode(ctx context.Context, shortCode string) (models.URL, error)
	CreateURL(ctx context.Context, url models.URL, shortCode string) error
	GetAllURLs(ctx context.Context) ([]models.URL, error)
	DeleteURL(ctx context.Context, shortCode string) error
	UpdateURL(ctx context.Context, url models.URL, shortCode string) error
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}
