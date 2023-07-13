package handler

import (
	"net/http"
)

func (handler *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Create URL"))

}

func (handler *Handler) GetURLByShortCode(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) GetAllURLs(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) DeleteURL(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) UpdateURL(w http.ResponseWriter, r *http.Request) {}
