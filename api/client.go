// Package api base client definition to use throughout
// the REPL
package api

import (
	"net/http"
	"time"

	"github.com/deexth/pokedex/api/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
