package main

import (
	"time"

	"github.com/iostate/pokedexcli/internal/pokeapi"
	"github.com/iostate/pokedexcli/internal/pokedex"
)

func main() {
	config := &config{
		client:  pokeapi.NewClient(5*time.Second, 5*time.Second),
		pokedex: pokedex.NewPokedex(),
	}

	repl := &repl{
		config: config,
	}

	repl.Start()
}
