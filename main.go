package main

import (
	"time"

	"github.com/coderjcronin/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		ownedPokemon:  make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
