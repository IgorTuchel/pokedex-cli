package main

import (
	"time"

	"github.com/igortuchel/pokedex-cli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(10 * time.Second)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
