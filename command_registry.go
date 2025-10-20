package main

import (
	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokecache.Cache, *pokeapi.LocationConfig, *pokeapi.PokemonConfig, *pokeapi.PokemonDataConfig, []string) error
}

func getCommandRegistry() map[string]cliCommand {
	commandRegistry := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "lists the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "lists the previous 20 locations",
			callback:    commandMap,
		},
		"explore": {
			name:        "explore",
			description: "lists all the pokemon in the specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catches a pokemon in the explored location",
			callback:    commandCatch,
		},
	}
	return commandRegistry
}
