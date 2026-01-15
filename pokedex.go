package main

import "fmt"

func commandPokedex(conf *config, args ...string) error {
	if len(conf.pokemon) == 0 {
		fmt.Println("You have not caught any pokemon yet")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for p := range conf.pokemon {
		fmt.Printf(" - %s\n", p)
	}

	return nil
}
