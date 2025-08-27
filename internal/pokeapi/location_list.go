package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Client) ListLocations(pageURL *string) (*LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache
	cacheData, ok := api.cache.Get(url)
	if ok {
		var locationAreas LocationAreasResponse
		if err := json.Unmarshal(cacheData, &locationAreas); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cache data: %w", err)
		}
		return &locationAreas, nil
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

	var locationAreas LocationAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationAreas); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Update cache
	jsonData, err := json.Marshal(locationAreas)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response for caching: %w", err)
	}
	api.cache.Add(url, jsonData)

	return &locationAreas, nil
}
