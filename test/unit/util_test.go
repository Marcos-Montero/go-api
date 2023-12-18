package utils_test

import (
	"os"
	"testing"

	"MyGoProject/internal/utils"
)

// TestCheckErr tests the CheckErr function
func TestCheckErr(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("CheckErr did not panic on error")
        }
    }()

    // Trigger an error
    utils.CheckErr(&utils.MockError{})
}

// TestGetEnvWithFallback tests the GetEnvWithFallback function
func TestGetEnvWithFallback(t *testing.T) {
    const envKey = "TEST_ENV"
    const fallbackValue = "default"

    // Case 1: Environment variable not set
    if val := utils.GetEnvWithFallback(envKey, fallbackValue); val != fallbackValue {
        t.Errorf("Expected %s; got %s", fallbackValue, val)
    }

    // Case 2: Environment variable is set
    const expectedValue = "value"
    os.Setenv(envKey, expectedValue)
    if val := utils.GetEnvWithFallback(envKey, fallbackValue); val != expectedValue {
        t.Errorf("Expected %s; got %s", expectedValue, val)
    }
    os.Unsetenv(envKey)
}

// TestContainsString tests the ContainsString function
func TestContainsString(t *testing.T) {
    slice := []string{"apple", "banana", "orange"}

    // Case 1: String is in the slice
    if !utils.ContainsString(slice, "banana") {
        t.Errorf("Expected true; got false")
    }

    // Case 2: String is not in the slice
    if utils.ContainsString(slice, "grape") {
        t.Errorf("Expected false; got true")
    }
}

// MockError is a mock error type for testing
type MockError struct{}

func (m *MockError) Error() string {
    return "mock error"
}
