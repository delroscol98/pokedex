package main

import (
	"fmt"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func commandPokedex(pokedex map[string]pokeapi.PokemonDataConfig, cache *pokecache.Cache, locConfig *pokeapi.LocationConfig, pokemonConfig *pokeapi.PokemonConfig, pokemonDataConfig *pokeapi.PokemonDataConfig, args []string) error {
	fmt.Println("Your Pokedex:")

	for pokemon := range pokedex {
		fmt.Printf("- %s\n", pokemon)
	}

	return nil
}
