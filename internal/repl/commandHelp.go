package repl

import "fmt"

func commandHelp(c *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	for _, c := range getCommands() {
		fmt.Printf("  %s - %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}
