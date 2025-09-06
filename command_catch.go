package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *config, arg *string) error {
	pokemon, err := c.pokeapiClient.Catch(arg)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at " + *arg + "...")
	chance := pokemon.Base_experience - rand.IntN(pokemon.Base_experience)

	if chance < 100 {
		fmt.Println(*arg + " was caught!")
		c.pokedex.Add(pokemon)
	} else {
		fmt.Println(*arg + " escaped!")
	}

	// fmt.Printf("Exploring %v...\n", *arg)
	// fmt.Println("Found Pokemon:")
	// fmt.Printf("Name: %v\n", pokemon.Name)
	// fmt.Printf("XP: %v\n", pokemon.Base_experience)
	// for _, mons := range pokemon.Stats {
	// 	fmt.Println("- " + mons.Stat.Name)

	// }

	return nil
}