package main

import "fmt"

func commandMapb(cfg *config, args ...string) error {
	locationAreas, err := cfg.client.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		fmt.Print("error running map command")
		return err
	}

	cfg.nextLocationsURL = locationAreas.Next
	if locationAreas.Previous != nil {
		cfg.previousLocationsURL = locationAreas.Previous
	}

	printLocationAreaResults(locationAreas)

	return nil
}
