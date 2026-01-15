package main

import (
	"errors"
	"fmt"
)

func commandMap(conf *config) error {
	if conf.next == nil && conf.previous != nil {
		return errors.New("you are on the last page")
	}

	locations, err := conf.pokeAPI.LocationAreas(conf.next)
	if err != nil {
		return err
	}

	conf.next = locations.Next
	conf.previous = locations.Previous

	for _, location := range locations.Results {

		fmt.Printf("%s \n", location.Name)
	}

	return nil
}

func commandMapB(conf *config) error {
	if conf.previous == nil {
		return errors.New("you are on the first page")
	}

	locations, err := conf.pokeAPI.LocationAreas(conf.previous)
	if err != nil {
		return err
	}

	conf.next = locations.Next
	conf.previous = locations.Previous

	for _, location := range locations.Results {

		fmt.Printf("%s \n", location.Name)
	}

	return nil
}
