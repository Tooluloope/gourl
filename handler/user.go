package handler

import (
	"fmt"
	"net/http"

	"github.com/Tooluloope/gourl/models"
)

type AuthService interface {
	AuthenticateUser(username, password string) (models.User, error)
	RegisterUser(user models.User) error
}

func (handler *Handler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	handler.Service.RegisterUser(r.Context(), models.User{})
	fmt.Println("RegisterUser Handler")

}
