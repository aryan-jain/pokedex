package main

import (
	"internal/cache"
	"internal/poke_api"
	"internal/repl"
	"time"
)

func main() {
	poke_cache := cache.NewCache(100)
	client := poke_api.NewClient(time.Second*5, poke_cache)
	cfg := repl.Config{
		Client: client,
	}

	repl.StartRepl(&cfg)
}
