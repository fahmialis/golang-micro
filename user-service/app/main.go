package main

import (
	"log"
	"net/http"
	"os"

	"user-service/internal/config"
	"user-service/internal/repository"
	"user-service/internal/router"
	"user-service/internal/service"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:55000"
	}

	client := config.NewMongoClient(mongoURI)
	db := client.Database("local_mongo_db")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	handler := router.New(userService)

	log.Println("User service running on :4001")
	http.ListenAndServe(":4001", handler)
}
