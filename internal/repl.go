package internal

import (
	"bufio"
	"errors"
	"fmt"
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
		return "", nil, errors.New("Nothing was typed in, please enter a command. Enter the command 'help' to get information about the app.")
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
	intro()
	printRepl()
	for reader.Scan() {
		text, flags, err := cleanInput(reader.Text())

		if err != nil {
			fmt.Println(err)
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

func intro() {
	fmt.Println(`
_______  ________  ________  ___  ________   ________          ________  ___       ___     
|\   __  \|\   __  \|\   ____\|\  \|\   ___  \|\   ____\        |\   ____\|\  \     |\  \    
\ \  \|\  \ \  \|\  \ \  \___|\ \  \ \  \\ \  \ \  \___|        \ \  \___|\ \  \    \ \  \   
 \ \   _  _\ \   __  \ \  \    \ \  \ \  \\ \  \ \  \  ___       \ \  \    \ \  \    \ \  \  
  \ \  \\  \\ \  \ \  \ \  \____\ \  \ \  \\ \  \ \  \|\  \       \ \  \____\ \  \____\ \  \ 
   \ \__\\ _\\ \__\ \__\ \_______\ \__\ \__\\ \__\ \_______\       \ \_______\ \_______\ \__\
    \|__|\|__|\|__|\|__|\|_______|\|__|\|__| \|__|\|_______|        \|_______|\|_______|\|__|`)
	fmt.Println("Welcome to the Racing CLI, here you can find the latest race results!")
	fmt.Println(`
We currently support te following racing sports
	- Formula 1`)
	fmt.Println("To get started, enter the name of the sport you want to get results from.")
	fmt.Println("NOTE: If a sport has number in it, for example 'Formula 1' enter it as follows: FormulaOne")
}
