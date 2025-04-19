package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	maxExperience := 500                                                  // A reasonable cap for base_experience
	probability := 100 - ((pokemon.BaseExperience * 100) / maxExperience) // Scaled to a range of 0-100

	if rand.Intn(100) < probability {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.ownedPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
