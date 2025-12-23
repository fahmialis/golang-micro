package router

import (
	"net/http"

	"gateway-service/internal/proxy"
)

func New() http.Handler {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("gateway ok"))
	})

	// User service routes
	mux.Handle("/users/", proxy.User())
	mux.Handle("/users", proxy.User())

	return mux
}
