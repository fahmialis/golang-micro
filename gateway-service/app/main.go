package main

import (
	"log"

	"gateway-service/internal/server"
)

func main() {
	log.Println("🚀 Starting gateway service")
	server.Run()
}
