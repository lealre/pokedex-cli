package commands

import (
	"fmt"
	"io"
	"net/http"
)

func commandMapBack(config *Config) error {

	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	previousUrl := config.Previous

	if value, ok := config.Cache.Get(previousUrl); ok {
		fmt.Printf("Using cahce for %s\n", previousUrl)
		return printLocations(value, config)
	}

	// Request
	res, err := http.Get(previousUrl)
	if err != nil {
		return fmt.Errorf("error occurred: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("error occurred: %w", err)
	}

	// Add cache
	config.Cache.Add(previousUrl, body)

	return printLocations(body, config)
}
