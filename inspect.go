package main

import (
	"errors"
	"fmt"
)

func commandInspect(conf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("pokemon name not given")
	}

	mapName, ok := conf.pokemon[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon yet")
		return nil
	}

	fmt.Printf("Name: %s\n", args[0])
	fmt.Printf("Height: %d\n", mapName.Height)
	fmt.Printf("Weight: %d\n", mapName.Weight)
	fmt.Println("stats:")

	for _, stat := range mapName.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, t := range mapName.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
