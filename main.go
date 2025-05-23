package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lealre/pokedex-cli/commands"
	"github.com/lealre/pokedex-cli/pokecache"
)

func main() {

	// Set in-memory cache cleaning interval
	var cacheInterval time.Duration = 60
	cache := pokecache.NewCache(time.Second * cacheInterval)

	config := &commands.Config{Cache: cache}
	storage := &commands.Storage{Storage: make(map[string]commands.PokemonExperience)}
	cmds := commands.GetCommands(config, storage)

	for {
		fmt.Print("\nPokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		inputUser := scanner.Text()
		inputCleaned := cleanInput(inputUser)
		if len(inputCleaned) == 0 {
			cmds["exit"].Callback("")
		}
		baseCommand := inputCleaned[0]

		var argument string
		if len(inputCleaned) > 1 {
			argument = inputCleaned[1]
		}

		if command, ok := cmds[baseCommand]; ok {
			err := command.Callback(argument)
			if err != nil {
				fmt.Println(err)
			}
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
