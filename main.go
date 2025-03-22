package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lealre/pokedex-cli/commands"
)

func main() {

	config := &commands.Config{}
	cmds := commands.GetCommands(config)

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		inputUser := scanner.Text()
		inputCleaned := cleanInput(inputUser)
		if len(inputCleaned) == 0 {
			cmds["exit"].Callback()
		}
		userCommand := inputCleaned[0]

		if command, ok := cmds[userCommand]; ok {
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
