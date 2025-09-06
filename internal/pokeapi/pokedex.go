package pokeapi

import (
	"errors"
	"fmt"
)

type Pokedex struct {
	Pokemons map[string][]Pokemon
}

func (p *Pokedex) Add (pokemon *Pokemon) error {
	if pokemon == nil{
		return errors.New("invalid pokemon")
	}

	p.Pokemons[pokemon.Name] = append(p.Pokemons[pokemon.Name], *pokemon)
	return nil
}

func (p *Pokedex) Inspect(name string) error {
	pokemons, ok := p.Pokemons[name]
	if !ok {
		return errors.New("cannot find the pokemon(s)")
	}
	for count, pokemon := range pokemons {
		fmt.Printf("[%v] Name: %v\n",count,pokemon.Name)
		fmt.Printf("Height: %v\n",pokemon.Height)
		fmt.Printf("Weight: %v\n",pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats{
			fmt.Printf("-%v: %v\n",stat.Stat.Name,stat.Effort)
		}
		
		fmt.Println("Types:")
		for _, poketype := range pokemon.Types{
			fmt.Printf("- %v\n",poketype.Type.Name)
		}
		fmt.Println("")
	}
	return nil
}

func (p *Pokedex) PokeDex() error {
	fmt.Println("My PokeDex:")
	for name, pokemon := range p.Pokemons {
		fmt.Printf("- [%v] %v\n",len(pokemon),name)
	}
	return nil
}