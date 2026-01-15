package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(conf *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if (commandName == "explore" || commandName == "catch") && len(words) == 2 {
			conf.name = &words[1]
		}

		availableCommands := getCommands()

		cmmd, ok := availableCommands[commandName]
		if !ok {
			fmt.Print("Unknown command\n")
			continue
		}

		if err := cmmd.callback(conf); err != nil {
			fmt.Printf("something went wrong: %v", err)
		}

	}
}

func cleanInput(text string) []string {
	newText := strings.ToLower(text)
	return strings.Fields(newText)
}
