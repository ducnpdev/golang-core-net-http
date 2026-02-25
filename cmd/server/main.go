package main

import (
	"log"
	"os"
	"time"

	"github.com/ducnpdev/golang-core-net-http/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := server.New(":" + port)

	log.Printf("Starting server on %s", port)

	if err := srv.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}

	time.Sleep(100 * time.Millisecond)
	log.Println("Server stopped cleanly")
}
