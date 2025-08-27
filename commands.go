package main

import (
	"fmt"
	"os"
)

// Represents a CLI command
type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

// Generate command directory
func getCommandDirectory() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 location areas at a time",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 location areas at a time",
			callback:    commandMapb,
		},
	}
}

// Exit the program
func commandExit(cfg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

// Display usage information
func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommandDirectory() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

// Maps command (move forward in the results)
func commandMap(cfg *config) error {
	locationAreas, err := cfg.client.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		fmt.Print("error running map command")
		return err
	}

	cfg.nextLocationsURL = locationAreas.Next
	if locationAreas.Previous != nil {
		cfg.previousLocationsURL = locationAreas.Previous
	}

	for _, result := range locationAreas.Results {
		fmt.Printf("%s\n", result.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	locationAreas, err := cfg.client.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		fmt.Print("error running map command")
		return err
	}

	cfg.nextLocationsURL = locationAreas.Next
	if locationAreas.Previous != nil {
		cfg.previousLocationsURL = locationAreas.Previous
	}

	for _, result := range locationAreas.Results {
		fmt.Printf("%s\n", result.Name)
	}
	return nil
}
