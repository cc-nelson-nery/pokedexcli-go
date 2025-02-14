package main

import "fmt"

func commandPokedex(cfg *Config) error {
	fmt.Println("Your Pokedex:")
	for k := range cfg.pokedex {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
