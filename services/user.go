package services

import (
	"context"
	"fmt"

	"github.com/Tooluloope/gourl/models"
)

func (service *Service) AuthenticateUser(ctx context.Context, username, password string) (models.User, error) {
	return models.User{}, nil
}

func (service *Service) RegisterUser(ctx context.Context, user models.User) error {
	fmt.Println("RegisterUser Service")

	service.store.RegisterUser(ctx, user)
	return nil
}
