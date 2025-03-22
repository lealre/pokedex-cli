package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lealre/pokedex-cli/commands"
)

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		inputUser := scanner.Text()
		inputCleaned := cleanInput(inputUser)
		if len(inputCleaned) == 0 {
			commands.GetCommands()["exit"].Callback()
		}
		userCommand := inputCleaned[0]

		if command, ok := commands.GetCommands()[userCommand]; ok {
			command.Callback()
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
