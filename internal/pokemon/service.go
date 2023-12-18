package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Pokemon represents the structure of a Pokémon
type Pokemon struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    // Add other relevant fields
}

// Service provides methods to interact with Pokémon data
type Service struct {
    // Dependencies, like a database or external API client, can be added here
}

// NewService creates a new instance of the Pokémon service
func NewService() *Service {
    return &Service{}
}

// FetchPokemon fetches Pokémon data by ID
func (s *Service) FetchPokemon(id int) (*Pokemon, error) {
    // For demonstration, fetching data from the external PokeAPI
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching Pokémon: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("received non-200 status code %d", resp.StatusCode)
    }

    var pokemon Pokemon
    err = json.NewDecoder(resp.Body).Decode(&pokemon)
    if err != nil {
        return nil, fmt.Errorf("error decoding Pokémon data: %v", err)
    }

    return &pokemon, nil
}
