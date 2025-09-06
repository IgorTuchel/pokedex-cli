package main

import (
	"time"

	"github.com/igortuchel/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeDex := pokeapi.Pokedex{
		Pokemons: map[string][]pokeapi.Pokemon{},
	}
	pokeClient := pokeapi.NewClient(10 * time.Second,10 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: &pokeDex,
	}

	startRepl(cfg)
}
