package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(conf *config) error {
	if conf.name == nil {
		return errors.New("no pokemon name provided")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *conf.name)

	poke, err := conf.pokeAPI.PokemonAPI(*conf.name)
	if err != nil {
		return err
	}

	w := rand.Intn(poke.BaseExperience)
	if (w + 1) > 40 {
		fmt.Printf("%s was caught!\n", *conf.name)
		conf.pokemon[*conf.name] = poke

		return nil
	}

	fmt.Printf("%s escaped!\n", *conf.name)

	return nil
}
