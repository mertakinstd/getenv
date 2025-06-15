package getenv_test

import (
	"os"
	"testing"

	getenv "github.com/mertakinstd/getenv"
)

func TestLoad_Default(t *testing.T) {
	// Clean environment before test
	os.Clearenv()

	// Run test
	err := getenv.Load(false).Default()
	if err != nil {
		t.Errorf("Load().Default() returned error: %v", err)
	}

	// Expected values
	tests := []struct {
		key      string
		expected string
	}{
		{"PORT", "8080"},
		{"IP", "10.0.0.1"},
		{"HOST", "10.0.0.1:8080"},
		{"ENV", "DEFAULT"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.expected {
				t.Errorf("%s = %v, want %v", tt.key, got, tt.expected)
			}
		})
	}
}

func TestLoad_NonExistentFile(t *testing.T) {
	// Clean environment before test
	os.Clearenv()

	// Test in production mode (.env.production file doesn't exist)
	err := getenv.Load(false).Production()
	if err != nil {
		t.Errorf("Should not return error for non-existent file: %v", err)
	}
}

func TestLoad_DuplicateEnvironmentVariables(t *testing.T) {
	// Clean environment before test
	os.Clearenv()

	// Set initial value
	os.Setenv("PORT", "3000")

	// Load .env file
	err := getenv.Load(false).Default()
	if err != nil {
		t.Errorf("Load().Default() returned error: %v", err)
	}

	// Verify that the original value is preserved
	if got := os.Getenv("PORT"); got != "3000" {
		t.Errorf("PORT = %v, want 3000", got)
	}
}

func TestLoad_UpdateEnvironmentVariables(t *testing.T) {
	// Clean environment before test
	os.Clearenv()

	// Set initial value
	os.Setenv("PORT", "3000")

	// Load .env file
	err := getenv.Load(true).Default()
	if err != nil {
		t.Errorf("Load().Default() returned error: %v", err)
	}

	// Verify that the value is updated
	if got := os.Getenv("PORT"); got != "8080" {
		t.Errorf("PORT = %v, want 8080", got)
	}
}

func TestLoad_Development(t *testing.T) {
	// Clean environment before test
	os.Clearenv()

	// Run test
	err := getenv.Load(false).Development()
	if err != nil {
		t.Errorf("Load().Development() returned error: %v", err)
	}

	// Expected values
	tests := []struct {
		key      string
		expected string
	}{
		{"PORT", "8080"},
		{"IP", "10.0.0.1"},
		{"HOST", "10.0.0.1:8080"},
		{"ENV", "DEV"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.expected {
				t.Errorf("%s = %v, want %v", tt.key, got, tt.expected)
			}
		})
	}
}

func TestLoad_Production(t *testing.T) {
	// Clean environment before test
	os.Clearenv()

	// Run test
	err := getenv.Load(false).Production()
	if err != nil {
		t.Errorf("Load().Production() returned error: %v", err)
	}

	// Expected values
	tests := []struct {
		key      string
		expected string
	}{
		{"PORT", "8080"},
		{"IP", "10.0.0.1"},
		{"HOST", "10.0.0.1:8080"},
		{"ENV", "PROD"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := os.Getenv(tt.key); got != tt.expected {
				t.Errorf("%s = %v, want %v", tt.key, got, tt.expected)
			}
		})
	}
}
