package main

import (
	"errors"
)

func commandInspect(c *config, arg *string) error {
	if arg == nil {
		return errors.New("usage inspect <pokemon name>")
	}
	err := c.pokedex.Inspect(*arg)
	if err != nil {
		return err
	}
	
	return nil
}