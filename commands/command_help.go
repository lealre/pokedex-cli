package commands

import (
	"fmt"
)

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, cmd := range GetCommands(config) {
		fmt.Printf("%s: %s\n", key, cmd.Description)
	}
	return nil
}
