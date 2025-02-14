package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", *cfg.arg)

	if _, ok := cfg.pokedex[*cfg.arg]; ok {
		fmt.Printf("%s was caught!\n", *cfg.arg)
		return nil
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(cfg.arg)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience) + 30

	if pokemon.BaseExperience > chance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	}

	return nil
}
