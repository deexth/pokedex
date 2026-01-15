package main

import (
	"time"

	"github.com/deexth/pokedex/api"
)

func main() {
	client := api.NewClient(10*time.Second, 7*time.Minute)

	conf := &config{
		pokemon: map[string]api.Pokemon{},
		pokeAPI: client,
	}

	startRepl(conf)
}
