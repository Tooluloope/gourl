package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Tooluloope/gourl/server/middleware"
	"github.com/Tooluloope/gourl/server/services"
	"github.com/gorilla/mux"
)

type Handler struct {
	Service *services.Service
	Router  *mux.Router
	Server  *http.Server
}

func NewHandler(service *services.Service) *Handler {

	handler := &Handler{
		Service: service,
	}
	handler.Router = mux.NewRouter()
	handler.mapRoutes()
	handler.Router.Use(middleware.JSONMiddleware)
	port := os.Getenv("PORT")
	handler.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler.Router,
	}

	return handler
}

func (handler *Handler) mapRoutes() {
	handler.Router.HandleFunc("/api/v1/createuser", handler.RegisterUser).Methods("POST")
	handler.Router.HandleFunc("/api/v1/login", handler.AuthenticateUser).Methods("POST")
	handler.Router.HandleFunc("/api/v1/createurl", middleware.JWTAuth(handler.CreateURL)).Methods("POST")
	handler.Router.HandleFunc("/api/v1/geturl", middleware.JWTAuth(handler.GetURLByShortCode)).Methods("GET")
	handler.Router.HandleFunc("/api/v1/getallurls", middleware.JWTAuth(handler.GetAllURLs)).Methods("GET")
	handler.Router.HandleFunc("/api/v1/deleteurl/{id}", middleware.JWTAuth(handler.DeleteURL)).Methods("DELETE")
	handler.Router.HandleFunc("/api/v1/updateurl", middleware.JWTAuth(handler.UpdateURL)).Methods("PUT")
	handler.Router.HandleFunc("/r/{shortCode}", handler.RedirectToShortURL).Methods("GET")
}

func (handler *Handler) RunServer() {

	go func() {
		fmt.Printf("Starting server on %s", handler.Server.Addr)
		if err := handler.Server.ListenAndServe(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	defer handler.StopServer()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")

}

func (handler *Handler) StopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := handler.Server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
