package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cc-nelson-nery/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	arg           *string
	pokedex       map[string]pokeapi.Pokemon
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputs := cleanInput(scanner.Text())
		if len(inputs) == 0 {
			continue
		}

		command, ok := getCommands()[inputs[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if len(inputs) > 1 {
			arg := inputs[1]
			config.arg = &arg
		}

		err := command.callback(config)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "",
			callback:    commandPokedex,
		},
	}
}
