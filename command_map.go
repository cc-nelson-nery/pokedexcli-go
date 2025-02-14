package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}
