package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Tooluloope/gourl/models"
	"github.com/Tooluloope/gourl/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type PostURL struct {
	ShortCode   string `json:"shortCode" validate:"required"`
	OriginalURL string `json:"originalUrl" validate:"required,url"`
}

func convertToURL(postURL PostURL) models.URL {

	return models.URL{
		ShortCode:   postURL.ShortCode,
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

	queryParams := r.URL.Query()

	// Get a specific query parameter
	shortCode := queryParams.Get("shortCode")

	if shortCode == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, errors.New("shortCode is required"))
		return
	}

	url, err := handler.Service.GetURLByShortCode(r.Context(), shortCode)

	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, url)
}

func (handler *Handler) GetAllURLs(w http.ResponseWriter, r *http.Request) {

	urls, err := handler.Service.GetAllURLs(r.Context())

	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, urls)
}

func (handler *Handler) DeleteURL(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	urlId := vars["id"]
	if urlId == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	err := handler.Service.DeleteURL(r.Context(), urlId)

	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)

}

func (handler *Handler) UpdateURL(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) GetURLByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	urlId := vars["id"]
	if urlId == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	url, err := handler.Service.GetURLByShortCode(r.Context(), urlId)

	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, url)
}
