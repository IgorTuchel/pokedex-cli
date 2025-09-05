package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, value := range getCommands() {
		fmt.Printf("%v: %v\n",value.name,value.description)
	}
	return nil
}