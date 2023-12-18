package client

import (
	"MyGoProject/internal/pokemon"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// PokeAPIClient is a client for interacting with the PokeAPI
type PokeAPIClient struct {
    baseURL string
    httpClient *http.Client
}

// NewPokeAPIClient creates a new instance of PokeAPIClient
func NewPokeAPIClient(baseURL string) *PokeAPIClient {
    return &PokeAPIClient{
        baseURL: baseURL,
        httpClient: &http.Client{
            Timeout: time.Second * 10,
        },
    }
}

// GetPokemon fetches a Pokemon by its ID
func (client *PokeAPIClient) GetPokemon(id int) (*pokemon.Pokemon, error) {
    url := fmt.Sprintf("%s/pokemon/%d", client.baseURL, id)

    resp, err := client.httpClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making request to PokeAPI: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("received non-200 status code from PokeAPI: %d", resp.StatusCode)
    }

    var pkmn pokemon.Pokemon
    err = json.NewDecoder(resp.Body).Decode(&pkmn)
    if err != nil {
        return nil, fmt.Errorf("error decoding response from PokeAPI: %v", err)
    }

    return &pkmn, nil
}
