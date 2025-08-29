package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/iostate/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("command explore requires more one argument")
	}

	pokemonToCatch := args[0]

	fmt.Printf("Throwing a Pokeball at %s...", pokemonToCatch)

	pokemon, err := cfg.client.GetPokemonInfo(pokemonToCatch)
	if err != nil {
		return fmt.Errorf("error getting pokemon info: %w", err)
	}

	caught, err := attemptCatchPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("erorr attempting to catch pokemon: %w", err)
	}

	if caught {
		fmt.Printf("%s was caught and added to pokedex!\n", pokemon.Name)
		// Add to map
		cfg.pokedex.Add(pokemon)
	} else {
		fmt.Printf("%s was not caught!\n", pokemon.Name)
	}

	return nil
}

// Catches a Pokemon using chance
func attemptCatchPokemon(pokemon *pokeapi.Pokemon) (bool, error) {

	baseExperience := pokemon.BaseExperience

	catchChance := math.Max(10, math.Min(90, 100-float64(baseExperience)))

	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(100) + 1

	return float64(roll) <= catchChance, nil

}
