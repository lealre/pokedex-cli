package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		inputUser := scanner.Text()
		inputCleaned := cleanInput(inputUser)
		if len(inputCleaned) == 0 {
			getCommands()["exit"].callback()
		}
		userCommand := inputCleaned[0]

		if command, ok := getCommands()[userCommand]; ok {
			command.callback()
			continue
		}

		fmt.Println("Unknown command")
	}
}

func cleanInput(text string) []string {
	splited := strings.Fields(text)
	for i, word := range splited {
		splited[i] = strings.ToLower(word)
	}
	return splited
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", key, cmd.description)
	}
	return nil
}
