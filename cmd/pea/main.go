package main

import (
	"log"

	"github.com/jcmrs/gemini-prompt-engineer/internal/server"
)

func main() {
	srv := server.NewServer(":8080")
	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
