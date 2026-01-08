package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"user-service/internal/handler"
	"user-service/internal/service"
)

func New(userService *service.UserService) http.Handler {
	r := chi.NewRouter()

	// Middlewares (recommended)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userHandler := handler.NewUserHandler(userService)

	// Routes
	r.Get("/", userHandler.GetUsers)
	r.Get("/{id}", userHandler.GetUserByID)

	return r
}
