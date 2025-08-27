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
		userInput := lowercaseAndBreakUpWords(scanner.Text())

		if len(userInput) == 0 {
			fmt.Println()
			continue
		}

		userCommand := userInput[0]
		args := userInput[1:]
		if command, exists := commandDirectory[userCommand]; exists {
			if err := command.callback(r.config, args...); err != nil {
				fmt.Printf("error running command %s: %v\n", command.name, err)
			}
		} else {
			fmt.Printf("Unknown command: %s. Type 'help' for available commands", command.name)
		}
	}
}
