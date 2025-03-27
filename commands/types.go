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

type Storage struct {
	Storage map[string]PokemonExperience
}

// API JSON response structs
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

type PokemonExperience struct {
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonType struct {
	Slot int      `json:"slot"`
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
