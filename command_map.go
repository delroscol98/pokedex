package main

import (
	"errors"
	"fmt"

	"github.com/delroscol98/pokedex/pokeapi"
)

func commandMap(config *pokeapi.Config, args []string) error {
	var url string
	if args[0] == "mapb" && config.Previous == nil {
		return errors.New("You're on the first page")
	}

	if args[0] == "mapb" && config.Previous != nil {
		url = *config.Previous
	}

	if args[0] == "map" && config.Next == nil && config.Previous == nil {
		url = pokeapi.BaseUrl + "location-area/"
	}

	if args[0] == "map" && config.Next != nil {
		url = *config.Next
	}

	if err := pokeapi.GetLocationAreaAPI(config, url); err != nil {
		return err
	}

	for _, item := range config.Results {
		fmt.Println(item.Name)
	}

	return nil
}
