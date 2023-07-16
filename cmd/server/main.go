package main

import (
	"fmt"
	"log"

	"github.com/Tooluloope/gourl/server/database"
	"github.com/Tooluloope/gourl/server/handler"
	"github.com/Tooluloope/gourl/server/services"
	"github.com/joho/godotenv"
)

func RunServer() {
	db, err := handleDB()

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	service := services.NewService(db)

	httpHandler := handler.NewHandler(service)

	httpHandler.RunServer()
}

func handleDB() (*database.Database, error) {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	return database.Migrate(db)
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Environment variables loaded successfully!")

	RunServer()

}
