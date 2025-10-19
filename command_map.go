package main

import (
	"errors"
	"fmt"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func commandMap(cache *pokecache.Cache, locConfig *pokeapi.LocationConfig, pokemonConfig *pokeapi.PokemonConfig, args []string) error {
	var url string
	if args[0] == "mapb" && locConfig.Previous == nil {
		return errors.New("You're on the first page")
	}

	if args[0] == "mapb" && locConfig.Previous != nil {
		url = *locConfig.Previous
	}

	if args[0] == "map" && locConfig.Next == nil && locConfig.Previous == nil {
		url = pokeapi.BaseUrl + "location-area/?offset=0&limit=20"
	}

	if args[0] == "map" && locConfig.Next != nil {
		url = *locConfig.Next
	}

	if err := pokeapi.GetLocationAreaAPI(cache, locConfig, url); err != nil {
		return err
	}

	for _, item := range locConfig.Results {
		fmt.Println(item.Name)
	}

	return nil
}
