package main

import (
	"fmt"
	"os"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func commandExit(cache *pokecache.Cache, locConfig *pokeapi.LocationConfig, pokemonConfig *pokeapi.PokemonConfig, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
