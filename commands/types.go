package commands

import "github.com/lealre/pokedex-cli/pokecache"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(arg string) error
	Config      *Config
}

type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

type Location struct {
	Name string `json:"name"`
}

type Results struct {
	Results  []Location `json:"results"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
}

type PokemonEncounters struct {
	PokemonEncounters []Pokemon `json:"pokemon_encounters"`
}

type Pokemon struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}
