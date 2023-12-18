package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

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

func (s *Service) GetPokemonTypeCounts() (map[string]int, error) {
    const totalPokemon = 100 // Adjust the number based on your needs
    typeCounts := make(map[string]int)
    var wg sync.WaitGroup
    pokemonChannel := make(chan *Pokemon, totalPokemon)
    errorsChannel := make(chan error, totalPokemon)

    for i := 1; i <= totalPokemon; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            pkmn, err := s.FetchPokemon(id)
            if err != nil {
                errorsChannel <- err
                return
            }
            pokemonChannel <- pkmn
        }(i)
    }

    // Wait for all fetches to complete
    wg.Wait()
    close(pokemonChannel)
    close(errorsChannel)

    // Check for errors
    for err := range errorsChannel {
        if err != nil {
            return nil, err
        }
    }

    // Process the fetched Pokémon
    for pkmn := range pokemonChannel {
        for _, t := range pkmn.Types {
            typeCounts[t.Type.Name]++
        }
    }

    return typeCounts, nil
}