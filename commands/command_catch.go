package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

func commandCatch(pokemon string, config *Config, storage *Storage) error {

	url := "https://pokeapi.co/api/v2/pokemon/"
	fullUrl := url + pokemon

	if _, hasPokemon := storage.Storage[pokemon]; hasPokemon {
		fmt.Print("You already caught that Pokemon!")
		return nil
	}

	if value, ok := config.Cache.Get(fullUrl); ok {
		fmt.Printf("Using cahce for %s\n", fullUrl)
		return trhowPokeball(value, pokemon, storage)
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return fmt.Errorf("error occurred: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return fmt.Errorf("Pokemon %s not found", pokemon)
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	if err != nil {
		return fmt.Errorf("error occurred: %w", err)
	}

	config.Cache.Add(url, body)

	return trhowPokeball(body, pokemon, storage)

}

func trhowPokeball(val []byte, pokemon string, storage *Storage) error {
	maxExp := 608

	var pokemonExperience PokemonExperience
	err := json.Unmarshal(val, &pokemonExperience)
	if err != nil {
		return fmt.Errorf("error occurred getting information for %s: %w", pokemon, err)
	}

	captureChance := 1.0 - (float64(pokemonExperience.BaseExperience) / float64(maxExp))
	isCaptured := rand.Float64() < captureChance

	fmt.Printf("Throwing a Pokeball at %s...", pokemon)

	if isCaptured {
		fmt.Printf("%s was caught!\n", pokemon)
		storage.Storage[pokemon] = pokemonExperience
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon)
	return nil
}
