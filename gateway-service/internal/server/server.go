package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gateway-service/internal/router"
)

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router.New(),
	}

	// Start server
	go func() {
		log.Printf("🚪 Gateway running on :%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down gateway...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}

	log.Println("✅ Gateway exited properly")
}
