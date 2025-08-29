package main

import (
	"fmt"

	"github.com/iostate/pokedexcli/internal/pokeapi"
)

// Maps command (move forward in the results)
func commandMap(cfg *config, args ...string) error {
	locationAreas, err := cfg.client.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		fmt.Print("error running map command")
		return err
	}

	// Provide the next & previous set of location areas to be fetched
	cfg.nextLocationsURL = locationAreas.Next
	if locationAreas.Previous != nil {
		cfg.previousLocationsURL = locationAreas.Previous
	}

	printLocationAreaResults(locationAreas)

	return nil
}

func printLocationAreaResults(locationAreas *pokeapi.LocationAreasResponse) {
	for _, result := range locationAreas.Results {
		fmt.Printf("%s\n", result.Name)
	}
}
