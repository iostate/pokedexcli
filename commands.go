package main

// Represents a CLI command
type cliCommand struct {
	name        string
	description string
	// args are already "cleaned" (lowercase)
	callback func(cfg *config, args ...string) error
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
		"explore": {
			name:        "explore",
			description: "Explore a pokemon area by providing a location area as an argument",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon from pokedex",
			callback:    commandInspect,
		},
	}
}
