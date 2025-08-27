package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Client) GetPokemonInfo(pokemonName string) (*PokemonInfoResponse, error) {

	url := baseURL + "/pokemon/" + pokemonName

	if cacheData, ok := api.cache.Get(url); ok {
		var pokemonInfoResp PokemonInfoResponse
		if err := json.Unmarshal(cacheData, &pokemonInfoResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cache data: %w", err)
		}
		return &pokemonInfoResp, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var pokemonInfoResp PokemonInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&pokemonInfoResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	jsonData, err := json.Marshal(pokemonInfoResp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response for caching: %w", err)
	}

	api.cache.Add(url, jsonData)
	return &pokemonInfoResp, nil
}
