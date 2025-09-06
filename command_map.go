package main

import (
	"fmt"
)

func commandMap(c *config, _ *string) error {
	locationInfo, err := c.pokeapiClient.GetLocations(c.nextLocationUrl)
	if err != nil {
		return err
	}

	c.nextLocationUrl = locationInfo.Next
	c.prevLocationUrl = locationInfo.Previous

	for _, area := range locationInfo.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(c *config, _ *string) error {
	locationInfo, err := c.pokeapiClient.GetLocations(c.prevLocationUrl)
	if err != nil {
		return err
	}

	c.nextLocationUrl = locationInfo.Next
	c.prevLocationUrl = locationInfo.Previous

	for _, area := range locationInfo.Results {
		fmt.Println(area.Name)
	}

	return nil
}