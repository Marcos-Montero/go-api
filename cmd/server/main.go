package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MarcosMrod/go-api/internal/api"
	"github.com/MarcosMrod/go-api/pkg/config"
	"github.com/gorilla/mux"
)

func main() {
    // Load configurations
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Create a new router
    router := mux.NewRouter()

    // Initialize API handlers
    apiHandler := api.NewHandler()
    apiHandler.RegisterRoutes(router)

    // Configure the server
    srv := &http.Server{
        Handler:      router,
        Addr:         cfg.ServerAddress,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    // Start the server
    log.Printf("Starting server on %s", cfg.ServerAddress)
    if err := srv.ListenAndServe(); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
