package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(location string, config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	fullUrl := url + location

	// Request
	res, err := http.Get(fullUrl)
	if err != nil {
		return fmt.Errorf("error occurred requesting url %s: %w", fullUrl, err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("error occurred reading the response body: %w", err)
	}

	var PokemonEncounters PokemonEncounters
	err = json.Unmarshal(body, &PokemonEncounters)
	if err != nil {
		return fmt.Errorf("error occurred unmarshalling the response body: %w", err)
	}

	// Keep from here

	return nil
}
