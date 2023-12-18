package integration

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"MyGoProject/internal/api"
	"MyGoProject/internal/pokemon"

	"github.com/gorilla/mux"
)

func TestGetPokemonEndpoint(t *testing.T) {
    // Mock the PokemonService
    mockService := newMockPokemonService()

    // Initialize the API handler with the mocked service
    apiHandler := api.NewHandler(mockService)

    // Create a router and register routes
    router := mux.NewRouter()
    apiHandler.RegisterRoutes(router)

    // Create a test server using the router
    testServer := httptest.NewServer(router)
    defer testServer.Close()

    // Define the URL for the test server's endpoint
    url := fmt.Sprintf("%s/pokemon/1", testServer.URL)

    // Make a request to the test server
    response, err := http.Get(url)
    if err != nil {
        t.Fatalf("Failed to send request: %v", err)
    }
    defer response.Body.Close()

    // Check if the status code is what we expect
    if response.StatusCode != http.StatusOK {
        t.Errorf("Expected status OK; got %v", response.Status)
    }

    // Check the response body
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    expectedBody := `{"id":1,"name":"Bulbasaur"}`
    if string(body) != expectedBody {
        t.Errorf("Expected body %s; got %s", expectedBody, body)
    }
}

// newMockPokemonService returns a mocked version of the PokemonService
func newMockPokemonService() pokemon.Service {
    return &mockPokemonService{}
}

// mockPokemonService is a mock implementation of the PokemonService
type mockPokemonService struct{}

// FetchPokemon is a mock implementation of the FetchPokemon method
func (s *mockPokemonService) FetchPokemon(id int) (*pokemon.Pokemon, error) {
    return &pokemon.Pokemon{ID: 1, Name: "Bulbasaur"}, nil
}
