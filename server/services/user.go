package services

import (
	"context"

	"github.com/Tooluloope/gourl/server/models"
)

func (service *Service) AuthenticateUser(ctx context.Context, email, password string) (string, error) {

	user, err := service.store.AuthenticateUser(ctx, email)

	if err != nil {
		return "", err
	}

	if err := user.CheckPassword(password); err != nil {
		return "", err
	}

	return user.GenerateJWT()

}

func (service *Service) RegisterUser(ctx context.Context, user models.User) (models.User, error) {

	user, err := service.store.RegisterUser(ctx, user)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
