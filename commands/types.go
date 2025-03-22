package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
	Config      *Config
}

type Config struct {
	Next     string
	Previous string
}

type Location struct {
	Name string `json:"name"`
}

type Results struct {
	Results  []Location `json:"results"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
}
