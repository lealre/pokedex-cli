package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {

	var url string

	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = config.Next
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

	// Unmarshall
	var results Results
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println(err)
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
