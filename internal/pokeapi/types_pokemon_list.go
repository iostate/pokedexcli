package pokeapi

type PokemonRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}
type PokemonsFoundResponse struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
