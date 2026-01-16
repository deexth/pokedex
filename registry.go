package main

import "github.com/deexth/pokedex/api"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeAPI  api.Client
	next     *string
	previous *string
	pokemon  map[string]api.Pokemon
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in a Pokemon world",
			callback:    commandMap,
		},
		"clear": {
			name:        "clear",
			description: "Clear Pokedex screen",
			callback:    commandClear,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas in a Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemons in a specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "This command can be used to try catching a pokemon",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays your pokemons",
			callback:    commandPokedex,
		},
		"inspect": {
			name: "inspect",
			description: "Displays details about a pokemon",
			callback: commandInspect,
		},
	}
	return commands
}
