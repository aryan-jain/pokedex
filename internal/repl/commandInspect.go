package repl

import (
	"fmt"
)

func commandInspect(c *Config, args ...string) error {
	pokemon := args[0]
	pokemonDetails, ok := PokemonBelt[pokemon]
	if !ok {
		return fmt.Errorf("You have not yet caught a %s", pokemon)
	}
	fmt.Println("Name: ", pokemonDetails.Name)
	fmt.Println("Height: ", pokemonDetails.Height)
	fmt.Println("Weight: ", pokemonDetails.Weight)
	fmt.Println("Base Experience: ", pokemonDetails.BaseExperience)
	fmt.Println("Stats:")
	for _, s := range pokemonDetails.Stats {
		fmt.Println("  - ", s.Stat.Name, ": ", s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonDetails.Types {
		fmt.Println("  - ", t.Type.Name)
	}
	return nil
}
