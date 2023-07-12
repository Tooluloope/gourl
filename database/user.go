package database

import (
	"context"
	"fmt"

	"github.com/Tooluloope/gourl/models"
)

func (db *Database) AuthenticateUser(ctx context.Context, username, password string) (models.User, error) {
	return models.User{}, nil
}

func (db *Database) RegisterUser(ctx context.Context, user models.User) error {
	fmt.Println("RegisterUser DB")
	return nil
}
