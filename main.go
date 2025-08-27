package main

import (
	"time"

	"github.com/iostate/pokedexcli/internal/pokeapi"
)

func main() {
	config := &config{
		client: pokeapi.NewClient(5*time.Second, 5*time.Second),
	}

	repl := &repl{
		config: config,
	}

	repl.Start()
}
