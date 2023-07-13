package services

import (
	"context"

	"github.com/Tooluloope/gourl/models"
)

func (service *Service) AuthenticateUser(ctx context.Context, username, password string) (models.User, error) {
	return models.User{}, nil
}

func (service *Service) RegisterUser(ctx context.Context, user models.User) (models.User, error) {

	user, err := service.store.RegisterUser(ctx, user)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
