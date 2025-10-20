package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func commandCatch(cache *pokecache.Cache, locationConfig *pokeapi.LocationConfig, pokemonConfig *pokeapi.PokemonConfig, pokemonDataConfig *pokeapi.PokemonDataConfig, args []string) error {
	if len(args) == 1 {
		return errors.New("Which pokemon are you catching? Example: 'catch squirtle'.")
	}

	pokemon := args[1]

	for _, item := range pokemonConfig.PokemonEncounters {
		if pokemon == item.Pokemon.Name {
			url := item.Pokemon.URL
			if err := pokeapi.GetPokemonDataAPI(cache, pokemonDataConfig, url); err != nil {
				return err
			}

			throwPokeball(pokemon, pokemonDataConfig)

			return nil
		}
	}

	return errors.New("Specified pokemon not in explored area. Use 'explore <location-area>' to view pokemon in specified area.'")
}

func throwPokeball(pokemon string, pokemonDataConfig *pokeapi.PokemonDataConfig) {
	fmt.Printf("Throwing a Pokeball at %s\n", pokemon)

	catchBaseline := math.Floor(0.6 * float64(pokemonDataConfig.BaseExperience))
	catchNumber := rand.Float64() * float64(pokemonDataConfig.BaseExperience)

	if catchNumber < catchBaseline {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		fmt.Printf("%s was caught!\n", pokemon)
	}
}
