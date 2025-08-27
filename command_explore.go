package main

import (
	"fmt"
	"strings"

	"github.com/iostate/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("command explore requires more one argument")
	}

	area := args[0]

	pokemonFoundResp, err := cfg.client.ListPokemon(area)
	if err != nil {
		return fmt.Errorf("error listing pokemon: %w", err)
	}

	printEncounters(pokemonFoundResp)

	return nil
}

func printEncounters(pokemonFoundResp *pokeapi.PokemonsFoundResponse) {
	for _, encounter := range pokemonFoundResp.PokemonEncounters {
		fmt.Printf("- %s\n", strings.ToLower(encounter.Pokemon.Name))
	}
}
