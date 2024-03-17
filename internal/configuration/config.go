package configuration

type Config struct {
	LastCommand Command
	LastFlags   map[string]string
}

type Command struct {
	Name           string
	Description    string
	AvailableFlags map[string]string
	Command        func(cfg *Config) error
}
