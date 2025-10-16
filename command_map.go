package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config, args []string) error {
	var url string
	if args[0] == "mapb" && config.Previous == nil {
		return errors.New("You're on the first page")
	}

	if args[0] == "mapb" && config.Previous != nil {
		url = *config.Previous
	}

	if args[0] == "map" && config.Next == nil && config.Previous == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	if args[0] == "map" && config.Next != nil {
		url = *config.Next
	}

	if err := getLocationArea(config, url); err != nil {
		return err
	}

	for _, item := range config.Results {
		fmt.Println(item.Name)
	}

	return nil
}

func getLocationArea(config *Config, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	return nil
}
