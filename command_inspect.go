package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.ownedPokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon yet")
		return nil
	}

	fmt.Print(pokemon)

	return nil
}
