package pokeapi

import (
	"net/http"
	"time"

	internal "github.com/iostate/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      *internal.Cache
	httpClient *http.Client
}

func NewClient(timeout, cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: internal.NewCache(cacheInterval),
	}
}
