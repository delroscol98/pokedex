package main

import (
	"fmt"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func commandHelp(cache *pokecache.Cache, locConfig *pokeapi.LocationConfig, pokemonConfig *pokeapi.PokemonConfig, pokemonDataConfig *pokeapi.PokemonDataConfig, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommandRegistry() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
