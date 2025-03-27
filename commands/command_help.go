package commands

import (
	"fmt"
)

func commandHelp(_ string, config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, cmd := range GetCommands(config, nil) {
		fmt.Printf("%s: %s\n", key, cmd.Description)
	}
	return nil
}
