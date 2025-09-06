package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/igortuchel/pokedex-cli/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		var args *string
		
		userMsg := cleanInput(scanner.Text())
		if len(userMsg) == 0 {
			continue
		} else if len(userMsg) > 1 {
			args = &userMsg[1]
		}
		
		commandName := userMsg[0]
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Print("Unknown command\n")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

type config struct {
	pokeapiClient 	 pokeapi.Client
	nextLocationUrl	 *string
	prevLocationUrl  *string
	pokedex 		 *pokeapi.Pokedex
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	textFields := strings.Fields(lowerText)
	return textFields
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name: 		 "help",
			description: "Displays all the commands",
			callback:    commandHelp,
		},
		"map": {
			name: 		 "map",
			description: "Displays the next map",
			callback: commandMap,
		},
		"mapb": {
			name: 		 "mapb",
			description: "Displays the last map",
			callback: commandMapb,
		},
		"explore": {
			name: 		"explore <area>",
			description: "Displays all the pokemon in the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon>",
			description: "catch",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon>",
			description: "inspect",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "show all pokemons",
			callback: commandPokedex,
		},
	}
}