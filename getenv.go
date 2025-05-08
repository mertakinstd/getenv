package getenv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Loader struct {
	Platform string
}

func Load() *Loader {
	return &Loader{}
}

func (l *Loader) Default() error {
	l.Platform = ".env"
	return l.load()
}

func (l *Loader) Development() error {
	l.Platform = ".env.development"
	return l.load()
}

func (l *Loader) Production() error {
	l.Platform = ".env.production"
	return l.load()
}

func (l *Loader) load() error {
	data, err := l.loadFile()
	if err != nil {
		return err
	}

	Lines := strings.Split(string(data), "\n")

	for i := range Lines {
		if Lines[i] == "" {
			continue
		}

		if strings.HasPrefix(Lines[i], "#") {
			continue
		}

		kv := strings.SplitN(Lines[i], "=", 2)

		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])
		value = strings.Trim(value, `"'`)
		value = os.ExpandEnv(value)

		lookup, exists := os.LookupEnv(key)
		if exists && lookup != "" {
			fmt.Printf("Environment variable %s already set to %s, skipping\n", key, lookup)
			continue
		}

		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("failed to set environment variable: %s: %w", key, err)
		}
	}

	return nil
}

func (l *Loader) loadFile() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}

	path := filepath.Join(wd, l.Platform)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", fmt.Errorf("failed to read .env file")
	}

	return string(data), nil
}
