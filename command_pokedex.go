package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {

	// Get all pokemon
	allPokemon := cfg.pokedex.GetAllPokemon()

	if len(allPokemon) == 0 {
		fmt.Printf("Pokedex is empty\n")
		return nil
	}

	fmt.Printf("Your Pokedex: \n")
	for _, pokemon := range allPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
