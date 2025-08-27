package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Client) ListPokemon(area string) (*PokemonsFoundResponse, error) {
	url := baseURL + "/location-area/" + area

	// === Check Cache ===
	if cacheData, ok := api.cache.Get(url); ok {
		var pokemonFoundResp PokemonsFoundResponse
		if err := json.Unmarshal(cacheData, &pokemonFoundResp); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrFailedUnmarshal, err)
		}
		return &pokemonFoundResp, nil
	}

	// Cache miss
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var pokemonFoundResp PokemonsFoundResponse
	if err := json.NewDecoder(resp.Body).Decode(&pokemonFoundResp); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedDecode, err)
	}

	jsonData, err := json.Marshal(pokemonFoundResp)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedMarshal, err)
	}

	// Add to cache
	api.cache.Add(url, jsonData)

	return &pokemonFoundResp, nil
}
