package internal

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Woutjeee/racing_cli/internal/configuration"
)

func printRepl() {
	fmt.Print("$  ")
}

var flagRegex = regexp.MustCompile(`(-\S+)\s*('[^']*'|\S+)?`)

func cleanInput(text string) (string, map[string]string, error) {
	output := strings.TrimSpace(text)
	output = strings.ToLower(text)
	parts := strings.Fields(output)

	if output == "" {
		return "", nil, errors.New("Nothing was typed in, please enter a command.")
	}

	commandName := parts[0]
	matches := flagRegex.FindAllStringSubmatch(output, -1)

	result := make(map[string]string)
	for _, match := range matches {
		flag := match[1]
		value := match[2]

		if value != "" {
			value = strings.Trim(value, "'")
			result[flag] = value
		}
	}

	return commandName, result, nil
}

func StartReplLoop(cfg *configuration.Config) {
	commands := GetCommands(cfg)

	reader := bufio.NewScanner(os.Stdin)
	printRepl()
	for reader.Scan() {
		text, flags, err := cleanInput(reader.Text())

		if err != nil {
			log.Print(err)
			printRepl()
		} else {
			if command, exists := commands[text]; exists {
				cfg.LastCommand = command
				cfg.LastFlags = flags
				err := command.Command(cfg)

				if err != nil {
					// TODO: Print out command name here.
					fmt.Println("There was an error executing the command.", err)
				}
			} else {
				fmt.Println("Command did not exists")
			}
			printRepl()
		}
	}
}
