package main

import "fmt"

// Inspect a pokemon

func commandInspect(cfg *config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("command inspect requires more one argument")
	}

	pokemonToInspect := args[0]

	pokemon, err := cfg.pokedex.Get(pokemonToInspect)
	if err != nil {
		return err
	}

	pokemon.PrintStats()

	pokemon.PrintTypes()

	return nil
}
