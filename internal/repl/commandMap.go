package repl

import (
	"fmt"
)

func commandMap(cfg *Config, args ...string) error {
	loc, err := cfg.Client.ListLocationAreas(cfg.NextURL)
	if err != nil {
		return err
	}

	cfg.NextURL = loc.Next
	cfg.PrevURL = loc.Previous

	for _, l := range loc.Results {
		fmt.Println(l.Name)
	}
	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	loc, err := cfg.Client.ListLocationAreas(cfg.PrevURL)
	if err != nil {
		return err
	}

	cfg.NextURL = loc.Next
	cfg.PrevURL = loc.Previous

	for _, l := range loc.Results {
		fmt.Println(l.Name)
	}
	return nil
}
