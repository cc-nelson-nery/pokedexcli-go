package main

import (
	"fmt"
	"math/rand"
)

var pokedex = make(map[string]struct {
	Id             int
	Name           string
	BaseExperience int
})

func commandCatch(cfg *Config) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", *cfg.arg)
	pokemon, err := cfg.pokeapiClient.GetPokemon(cfg.arg)
	if err != nil {
		return err
	}

	if _, ok := pokedex[pokemon.Name]; ok {
		fmt.Printf("%s found in pokedex\n", pokemon.Name)
		return nil
	}

	chance := rand.Intn(pokemon.BaseExperience) + 20

	if pokemon.BaseExperience > chance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)

		pokedex[pokemon.Name] = struct {
			Id             int
			Name           string
			BaseExperience int
		}{
			Id:             pokemon.Id,
			Name:           pokemon.Name,
			BaseExperience: pokemon.BaseExperience,
		}
	}

	return nil
}
