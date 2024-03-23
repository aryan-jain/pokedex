package repl

import (
	"fmt"
	"time"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Invalid number of arguments. Usage: catch <pokemon name>")
	}
	pokemon := args[0]
	pokemonDetails, caught, err := c.Client.CatchPokemon(pokemon)
	fmt.Println("Threw a pokeball at ", pokemon)
	for range [3]int{} {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return err
	}
	if caught {
		fmt.Println(pokemon, " was caught!")
		PokemonBelt[pokemon] = pokemonDetails
	} else {
		fmt.Println(pokemon, " got away!")
	}
	return nil
}
