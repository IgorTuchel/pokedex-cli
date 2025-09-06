package main

import (
	"fmt"
)

func commandExplore(c *config, arg *string) error {
	pokemons, err := c.pokeapiClient.Explore(arg)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", *arg)
	fmt.Println("Found Pokemon:")
	for _, mons := range pokemons.Pokemon_encounters {
		fmt.Println("- " + mons.Pokemon.Name)
	}

	return nil
}