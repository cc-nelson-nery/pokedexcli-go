package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", *cfg.arg)
	pokemon, err := cfg.pokeapiClient.GetPokemon(cfg.arg)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience) + 20

	if pokemon.BaseExperience > chance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	}

	return nil
}
