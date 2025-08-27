package pokeapi

type PokemonInfoResponse struct {
	ID             int64
	Name           string
	BaseExperience int64 `json:"base_experience"`
}
