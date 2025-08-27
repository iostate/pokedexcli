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
			// very general error msg, does not specify
			// that we failed to unmarshal cache data.
			// consider changing this back
			// failed to unmarshal cache []bytes from cache
			// however, this error message is correct, in that we failed
			// to unmarshal something
			return nil, fmt.Errorf("%w: %v", ErrFailedUnmarshal, err)
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
		return nil, fmt.Errorf("%w: %v", ErrFailedDecode, err)
	}

	// Update cache
	jsonData, err := json.Marshal(locationAreas)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedMarshal, err)
	}
	api.cache.Add(url, jsonData)

	return &locationAreas, nil
}
