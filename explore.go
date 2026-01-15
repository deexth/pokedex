package main

import (
	"errors"
	"fmt"
)

func commandExplore(conf *config) error {
	if conf.name == nil {
		return errors.New("no pokemon name provided")
	}

	res, err := conf.pokeAPI.ExploreAPI(conf.name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *conf.name)
	fmt.Println("Found Pokemon:")

	for _, p := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}

	return nil
}
