package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Tooluloope/gourl/models"
	"github.com/Tooluloope/gourl/utils"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type RegisterUser struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,max=100,min=6"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=100,min=6"`
}

func convertToUser(user RegisterUser) models.User {
	return models.User{
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (handler *Handler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var user LoginUser

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	token, err := handler.Service.AuthenticateUser(r.Context(), user.Email, user.Password)

	if err != nil {
		utils.WriteJSONError(w, http.StatusUnauthorized, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, token)

}

func (handler *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user RegisterUser

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}
	userToCreate := convertToUser(user)

	createdUser, err := handler.Service.RegisterUser(r.Context(), userToCreate)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			utils.WriteJSONError(w, http.StatusConflict, fmt.Errorf("user with email %s already exists", user.Email))
			return
		}
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, createdUser)
}
