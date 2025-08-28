package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Client) GetPokemonInfo(pokemonName string) (*Pokemon, error) {

	url := baseURL + "/pokemon/" + pokemonName

	if cacheData, ok := api.cache.Get(url); ok {
		var pokemon Pokemon
		if err := json.Unmarshal(cacheData, &pokemon); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrFailedUnmarshal, err)
		}
		return &pokemon, nil
	}

	// Cache miss operations start here
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedDecode, err)
	}

	jsonData, err := json.Marshal(pokemon)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedMarshal, err)
	}

	// Add to cache
	api.cache.Add(url, jsonData)

	return &pokemon, nil
}
