package repl

import (
	"fmt"
	"os"
)

func commandExit(c *Config, args ...string) error {
	fmt.Println("Exitting...")
	os.Exit(0)
	return nil
}
