package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/delroscol98/pokedex/internal/pokeapi"
	"github.com/delroscol98/pokedex/internal/pokecache"
)

func main() {
	pokedex := map[string]pokeapi.PokemonDataConfig{}
	locConfig := &pokeapi.LocationConfig{}
	pokemonConfig := &pokeapi.PokemonConfig{}
	pokemonDataConfig := &pokeapi.PokemonDataConfig{}
	cache := pokecache.NewCache(5 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		var words []string
		if scanner.Scan() {
			words = cleanInput(scanner.Text())
			if len(words) == 0 {
				continue
			}
		}

		userInput := words[0]
		args := words

		command, exists := getCommandRegistry()[userInput]
		if exists {
			err := command.callback(pokedex, cache, locConfig, pokemonConfig, pokemonDataConfig, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.ToLower(strings.TrimSpace(text))

	if len(trimmedText) == 0 {
		return []string{}
	}

	words := strings.Split(trimmedText, " ")
	return words
}
