package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(_ string, config *Config) error {

	var url string

	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = config.Next
	}

	if value, ok := config.Cache.Get(url); ok {
		fmt.Printf("Using cahce for %s\n", url)
		return printLocations(value, config)
	}

	// Request
	res, err := http.Get(url)
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
	config.Cache.Add(url, body)

	return printLocations(body, config)
}

// Processes a byte slice representing JSON data
func printLocations(val []byte, config *Config) error {
	var results Results
	err := json.Unmarshal(val, &results)
	if err != nil {
		return fmt.Errorf("error occurred: %w", err)
	}

	for _, result := range results.Results {
		fmt.Printf("%s\n", result.Name)
	}

	config.Next = results.Next
	if results.Previous != nil {
		config.Previous = *results.Previous
	} else {
		config.Previous = ""
	}

	return nil
}
