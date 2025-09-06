package main

func commandPokedex(c *config, arg *string) error {
	c.pokedex.PokeDex()
	
	return nil
}