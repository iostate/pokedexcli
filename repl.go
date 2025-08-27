package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iostate/pokedexcli/internal/pokeapi"
)

type config struct {
	client               *pokeapi.Client
	previousLocationsURL *string
	nextLocationsURL     *string
}

type repl struct {
	config *config
}

func (r *repl) Start() {

	commandDirectory := getCommandDirectory()
	cli_header := "Welcome to the Pokedex!\n"
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s", cli_header)
	for {
		fmt.Print("Pokedex > ")
		// Run infinite loop as long as there's text to grab from os.Stdin

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %v\n", err)
			}
			fmt.Println("\nGoodbye!")
			break
		}
		userInput := scanner.Text()         // user's text
		cleanInput := cleanInput(userInput) // clean the input
		if len(cleanInput) == 0 {
			continue
		}
		fmt.Println()
		userCommand := cleanInput[0] // first word of clean input, aka command
		// check command registry and call callback, print any errors
		if command, exists := commandDirectory[userCommand]; exists {
			err := command.callback(r.config)
			if err != nil {
				fmt.Printf("error found: %v", err)
			}
		} else {
			fmt.Printf("Unknown command: %s. Type 'help' for available commands.\n", userCommand)
		}
	}
}
