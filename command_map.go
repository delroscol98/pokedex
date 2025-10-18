package main

import (
	"errors"
	"fmt"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func commandMap(cache *pokecache.Cache, config *pokeapi.Config, args []string) error {
	var url string
	if args[0] == "mapb" && config.Previous == nil {
		return errors.New("You're on the first page")
	}

	if args[0] == "mapb" && config.Previous != nil {
		url = *config.Previous
	}

	if args[0] == "map" && config.Next == nil && config.Previous == nil {
		url = pokeapi.BaseUrl + "location-area/?offset=0&limit=20"
	}

	if args[0] == "map" && config.Next != nil {
		url = *config.Next
	}

	if err := pokeapi.GetLocationAreaAPI(cache, config, url); err != nil {
		return err
	}

	for _, item := range config.Results {
		fmt.Println(item.Name)
	}

	return nil
}
