package main

import (
	"log"
	"net/http"
	"os"

	"user-service/internal/config"
	db "user-service/internal/db/gen"
	"user-service/internal/router"
	"user-service/internal/service"
)

func main() {
	postgresDSN := os.Getenv("POSTGRES_DSN")
	if postgresDSN == "" {
		postgresDSN = "postgres://admin:admin@localhost:6000/postgres?sslmode=disable"
	}

	dbPool := config.NewPostgresPool(postgresDSN)
	defer dbPool.Close()

	queries := db.New(dbPool)
	userService := service.NewUserService(queries)

	handler := router.New(userService)

	log.Println("User service running on :4001")
	http.ListenAndServe(":4001", handler)
}
