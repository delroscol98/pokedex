package main

import (
	"errors"
	"fmt"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func commandExplore(cache *pokecache.Cache, ocationConfig *pokeapi.LocationConfig, pokemonConfig *pokeapi.PokemonConfig, pokemonDataConfig *pokeapi.PokemonDataConfig, args []string) error {
	if len(args) == 1 {
		return errors.New("What location would you like to explore? Example: 'explore pastoria-city-area'.")
	}

	locationArea := args[1]
	fmt.Printf("Exploring %s...\n", locationArea)

	url := pokeapi.BaseUrl + "location-area/" + locationArea

	err := pokeapi.GetPokemonInLocationAreaAPI(cache, pokemonConfig, url)
	if err != nil {
		return errors.New("Invalid location specified. Try again!")
	}

	for _, item := range pokemonConfig.PokemonEncounters {
		pokemon := item.Pokemon.Name
		fmt.Println("-", pokemon)
	}

	return nil
}
