package main

import "fmt"

func commandExplore(cfg *Config) error {
	locations, err := cfg.pokeapiClient.LocationArea(cfg.arg)
	if err != nil {
		return err
	}

	for _, location := range locations.PokemonEncounters {
		fmt.Println(location.Pokemon.Name)
	}

	return nil
}
