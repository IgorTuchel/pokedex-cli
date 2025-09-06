package main

import (
	"fmt"
)

func commandHelp(_ *config, _ *string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range getCommands() {
		fmt.Printf("%v: %v\n",value.name,value.description)
	}
	return nil
}