package handler

import (
	"net/http"

	"github.com/Tooluloope/gourl/models"
)

type URLService interface {
	GetURLByShortCode(shortCode string) (models.URL, error)
	CreateURL(url models.URL) error
}

func (handler *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) GetURLByShortCode(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) GetAllURLs(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) DeleteURL(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) UpdateURL(w http.ResponseWriter, r *http.Request) {}
