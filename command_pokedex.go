package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.ownedPokemon) < 1 {
		fmt.Println("You haven't caught any pokemon yet!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.ownedPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
