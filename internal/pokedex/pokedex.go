package pokedex

import (
	"fmt"

	"github.com/iostate/pokedexcli/internal/pokeapi"
)

type Pokedex struct {
	caught map[string]*pokeapi.Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		caught: make(map[string]*pokeapi.Pokemon),
	}
}

func (p *Pokedex) Add(pokemon *pokeapi.Pokemon) {
	p.caught[pokemon.Name] = pokemon
}

func (p *Pokedex) Has(name string) bool {
	_, exists := p.caught[name]
	return exists
}

func (p *Pokedex) Remove(name string) {
	delete(p.caught, name)
}

func (p *Pokedex) Get(name string) (*pokeapi.Pokemon, error) {
	if !p.Has(name) {
		return nil, fmt.Errorf("%s was not found in pokedex", name)
	}
	return p.caught[name], nil
}

// Get all Pokemon in Pokedex
func (p *Pokedex) GetAllPokemon() []*pokeapi.Pokemon {
	var result []*pokeapi.Pokemon

	if len(p.caught) == 0 {
		return []*pokeapi.Pokemon{}
	}
	for _, entry := range p.caught {
		result = append(result, entry)
	}

	return result
}
