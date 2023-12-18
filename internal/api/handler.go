package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MarcosMRod/go-api/internal/pokemon"
	"github.com/gorilla/mux"
)

// Handler struct holds dependencies for the HTTP handlers
type Handler struct {
    PokemonService pokemon.Service
}

// NewHandler creates a new Handler with the given dependencies
func NewHandler(pokemonService pokemon.Service) *Handler {
    return &Handler{
        PokemonService: pokemonService,
    }
}

// RegisterRoutes sets up the routes for the HTTP server
func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/pokemon/{id}", h.GetPokemon).Methods("GET")
}

// GetPokemon handles the GET request for fetching Pokémon data
func (h *Handler) GetPokemon(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr, ok := vars["id"]
    if !ok {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid Pokémon ID", http.StatusBadRequest)
        return
    }

    pokemon, err := h.PokemonService.FetchPokemon(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResponse, err := json.Marshal(pokemon)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
