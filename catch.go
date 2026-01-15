package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	poke, err := conf.pokeAPI.PokemonAPI(args[0])
	if err != nil {
		return err
	}

	w := rand.Intn(poke.BaseExperience)
	if (w + 1) > 40 {
		fmt.Printf("%s was caught!\n", args[0])
		conf.pokemon[args[0]] = poke

		return nil
	}

	fmt.Printf("%s escaped!\n", args[0])

	return nil
}
