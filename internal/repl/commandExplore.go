package repl

import (
	"fmt"
)

func commandExplore(c *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Invalid number of arguments. Usage: explore <location name>")
	}
	loc := args[0]
	locationDetails, err := c.Client.ExploreLocation(loc)
	if err != nil {
		return err
	}

	fmt.Println("Exploring ", locationDetails.Name, "...")
	fmt.Println("Wild Pokemon Encounters:")
	for _, p := range locationDetails.PokemonEncounters {
		fmt.Println("  - ", p.Pokemon.Name)
	}
	return nil
}
