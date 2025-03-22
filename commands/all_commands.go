package commands

func GetCommands(config *Config) map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
			Config:      config,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func() error { return commandHelp(config) },
			Config:      config,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 locations",
			Callback:    func() error { return commandMap(config) },
			Config:      config,
		},
		"mapb": {
			Name:        "map",
			Description: "Displays 20 previous locations",
			Callback:    func() error { return commandMapBack(config) },
			Config:      config,
		},
	}
}
