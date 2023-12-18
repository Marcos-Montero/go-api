package utils

import (
	"log"
	"os"
)

// CheckErr logs the error and exits the program if the error is not nil.
func CheckErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

// GetEnvWithFallback tries to get an environment variable; if not found, it returns a fallback value.
func GetEnvWithFallback(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return fallback
    }
    return value
}

// ContainsString checks if a string is present in a slice.
func ContainsString(slice []string, str string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}
