package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MarcosMRod/go-api/internal/api"
	"github.com/MarcosMRod/go-api/internal/pokemon"
	"github.com/MarcosMRod/go-api/pkg/config"
	"github.com/gorilla/mux"
)

func main() {
    // Load configurations
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

		pokemonService := pokemon.NewService()
		// Initialize API handlers
		apiHandler := api.NewHandler(*pokemonService)
		// Create a new router
		router := mux.NewRouter()
		apiHandler.RegisterRoutes(router)
		router.HandleFunc("/pokemon/typecounts", apiHandler.GetPokemonTypeCounts).Methods("GET")


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
