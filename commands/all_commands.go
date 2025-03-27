package commands

func GetCommands(config *Config, storage *Storage) map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func(arg string) error { return commandExit(arg) },
			Config:      config,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func(arg string) error { return commandHelp(arg, config) },
			Config:      config,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 locations",
			Callback:    func(arg string) error { return commandMap(arg, config) },
			Config:      config,
		},
		"mapb": {
			Name:        "map",
			Description: "Displays 20 previous locations",
			Callback:    func(arg string) error { return commandMapBack(arg, config) },
			Config:      config,
		},
		"explore": {
			Name:        "explore",
			Description: "List all the Pokémon located in specific area",
			Callback:    func(arg string) error { return commandExplore(arg, config) },
			Config:      config,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a Pokemon",
			Callback:    func(arg string) error { return commandCatch(arg, config, storage) },
			Config:      config,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Check the stats of a Pokemon",
			Callback:    func(arg string) error { return commandInspect(arg, storage) },
			Config:      config,
		},
	}
}
