package repl

import "fmt"

func commandBag(c *Config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("Invalid number of arguments. Usage: bag")
	}
	fmt.Println("Pokemon in your bag:")
	for name, _ := range PokemonBelt {
		fmt.Println("  - ", name)
	}
	return nil
}
