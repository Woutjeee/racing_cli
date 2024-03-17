package internal

import (
	"fmt"
	"github.com/Woutjeee/racing_cli/internal/configuration"
)

func Help(cfg *configuration.Config) error {
	fmt.Println("Welcome to the help screen.")

	return nil
}

func exit() {

}

func GetCommands(cfg *configuration.Config) map[string]configuration.Command {
	return map[string]configuration.Command{
		"help": {
			Name:        "help",
			Description: "A command that shows the user how to get around the CLI",
			AvailableFlags: map[string]string{
				"-c": "List all the available commands",
			},
			Command: Help,
		},
	}
}
