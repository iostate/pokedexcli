package pokeapi

import (
	"errors"
	"net/http"
	"time"

	internal "github.com/iostate/pokedexcli/internal/pokecache"
)

var (
	ErrFailedDecode    = errors.New("failed to unmarshal cache data")
	ErrCacheCorrupted  = errors.New("cache corrupted")
	ErrFailedMarshal   = errors.New("failed to marshal data ")
	ErrFailedUnmarshal = errors.New("failed to unmarshal data")
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
