package config

import (
	"MyGoProject/internal/utils"
)

// Config represents the configuration settings for the application
type Config struct {
    ServerAddress string
    PokeAPIBaseURL string
    // Add other configuration fields as needed
}

// LoadConfig reads configuration settings and returns a Config struct
func LoadConfig() (*Config, error) {
    return &Config{
        ServerAddress: utils.GetEnvWithFallback("SERVER_ADDRESS", "localhost:8080"),
        PokeAPIBaseURL: utils.GetEnvWithFallback("POKEAPI_BASE_URL", "https://pokeapi.co/api/v2"),
        // Initialize other fields in a similar manner
    }, nil
}
