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

	pokemonInfoResp, err := cfg.client.GetPokemonInfo(pokemonToCatch)
	if err != nil {
		return fmt.Errorf("error getting pokemon info: %w", err)
	}

	caught, err := attemptCatchPokemon(pokemonInfoResp)
	if err != nil {
		return fmt.Errorf("erorr attempting to catch pokemon: %w", err)
	}

	if caught {
		fmt.Printf("%s was caught and added to pokedex!\n", pokemonToCatch)
		// Add to map
		cfg.caughtPokemon[pokemonToCatch] = pokemonInfoResp
	} else {
		fmt.Printf("%s was not caught!\n", pokemonToCatch)
	}

	return nil
}

func attemptCatchPokemon(pokemon *pokeapi.PokemonInfoResponse) (bool, error) {

	baseExperience := pokemon.BaseExperience

	catchChance := math.Max(10, math.Min(90, 100-float64(baseExperience)))

	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(100) + 1

	return float64(roll) <= catchChance, nil

}
