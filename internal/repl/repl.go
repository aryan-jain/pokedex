package repl

import (
	"bufio"
	"fmt"
	"internal/poke_api"
	"os"
	"strings"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	Client  poke_api.Client
	NextURL *string
	PrevURL *string
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Show help",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Show wild pokemon encounters in the location. [e.g. explore canalave-city-area]",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Catch a pokemon. [e.g. catch pikachu]",
			callback:    commandCatch,
		},
		"bag": {
			name:        "bag",
			description: "Show pokemon in your bag",
			callback:    commandBag,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Inspect a pokemon that you have caught. [e.g. inspect pikachu]",
			callback:    commandInspect,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func StartRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex> ")
		reader.Scan()
		textInput := reader.Text()

		words := cleanInput(textInput)
		if len(words) == 0 {
			continue
		}

		cmd := words[0]
		args := words[1:]

		c, ok := getCommands()[cmd]
		if !ok {
			fmt.Println("Command not found: ", cmd)
		} else {
			err := c.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
