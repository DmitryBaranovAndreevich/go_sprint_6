package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	router := server.CreateRouter(logger)
	err := router.Server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
