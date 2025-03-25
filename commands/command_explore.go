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

	if value, ok := config.Cache.Get(url); ok {
		fmt.Printf("Using cahce for %s\n", url)
		return printPokemons(value, location)
	}

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

	config.Cache.Add(url, body)

	return printPokemons(body, location)
}

func printPokemons(val []byte, location string) error {

	var pokemonEncounters PokemonEncounters
	err := json.Unmarshal(val, &pokemonEncounters)
	if err != nil {
		return fmt.Errorf("error occurred unmarshalling the response body: %w", err)
	}

	fmt.Printf("Exploring %s\n", location)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonEncounters.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil

}
