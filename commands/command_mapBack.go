package commands

import (
	"encoding/json"
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

	// Request
	res, err := http.Get(previousUrl)
	if err != nil {
		return fmt.Errorf("error occurred while requesting %s: %w", previousUrl, err)
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
