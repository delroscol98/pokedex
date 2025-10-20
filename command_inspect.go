package main

import (
	"errors"
	"fmt"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func commandInspect(pokedex map[string]pokeapi.PokemonDataConfig, cache *pokecache.Cache, locConfig *pokeapi.LocationConfig, pokemonConfig *pokeapi.PokemonConfig, pokemonDataConfig *pokeapi.PokemonDataConfig, args []string) error {
	if len(args) <= 1 {
		return errors.New("What pokemon would you like to inspect?")
	}

	pokemon := args[1]

	if pokemonData, exists := pokedex[pokemon]; !exists {
		fmt.Println("You have not caught that pokemon")
	} else {
		name := pokemonData.Name
		height := pokemonData.Height
		stats := pokemonData.Stats
		types := pokemonData.Types

		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Height: %d\n", height)
		fmt.Println("Stats:")
		for _, stat := range stats {
			fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, item := range types {
			fmt.Printf("- %s\n", item.Type.Name)
		}
	}

	return nil
}
