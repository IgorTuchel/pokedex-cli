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
		
		userMsg := cleanInput(scanner.Text())
		if len(userMsg) == 0 {
			continue
		}
		
		commandName := userMsg[0]
		
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg)
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
	callback    func(*config) error
}

type config struct {
	pokeapiClient 		 pokeapi.Client
	nextLocationUrl	 *string
	prevLocationUrl  *string
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
			description: "Displays the map",
			callback: commandMap,
		},
	}
}