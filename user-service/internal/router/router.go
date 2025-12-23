package router

import (
	"net/http"

	"user-service/internal/handler"
	"user-service/internal/service"
)

func New(userService *service.UserService) http.Handler {
	mux := http.NewServeMux()

	userHandler := handler.NewUserHandler(userService)

	mux.HandleFunc("/fetch", userHandler.GetUsers)
	mux.HandleFunc("/health", userHandler.Health)

	return mux
}
