package poke_api

import (
	"internal/cache"
	"net/http"
	"time"
)

type Client struct {
	client http.Client
	cache  *cache.Cache
}

func NewClient(timeout time.Duration, c *cache.Cache) Client {
	return Client{
		client: http.Client{
			Timeout: timeout,
		},
		cache: c,
	}
}
