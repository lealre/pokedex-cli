package commands

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", key, cmd.Description)
	}
	return nil
}
