package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tooluloope/gourl/models"
	"github.com/Tooluloope/gourl/utils"
	"github.com/go-playground/validator"
)

type PostURL struct {
	ShortURL    string `json:"shortUrl" validate:"required"`
	OriginalURL string `json:"originalUrl" validate:"required,url"`
}

func convertToURL(postURL PostURL) models.URL {

	return models.URL{
		ShortURL:    postURL.ShortURL,
		OriginalURL: postURL.OriginalURL,
	}
}

func (handler *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {

	var postURL PostURL

	// Decode the request body into struct and failed if any error occur
	if err := json.NewDecoder(r.Body).Decode(&postURL); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()

	if err := validate.Struct(postURL); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	// Create the URL
	url, err := handler.Service.CreateURL(r.Context(), convertToURL(postURL))

	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, url)
}

func (handler *Handler) GetURLByShortCode(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) GetAllURLs(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) DeleteURL(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) UpdateURL(w http.ResponseWriter, r *http.Request) {}
