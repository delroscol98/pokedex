package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/delroscol98/pokedex/internal/pokecache"
)

func handleGetRequest(cache *pokecache.Cache, url string) ([]byte, error) {
	var (
		data   []byte
		res    *http.Response
		err    error
		exists bool
	)

	data, exists = cache.Get(url)
	if !exists {
		res, err = http.Get(url)
		if err != nil {
			return nil, err
		}

		data, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("error reading response: %w", err)
		}
	}

	cache.Add(url, data)

	return data, nil
}

func GetLocationAreaAPI(cache *pokecache.Cache, locationConfig *LocationConfig, url string) error {
	data, err := handleGetRequest(cache, url)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &locationConfig); err != nil {
		return err
	}

	return nil
}

func GetPokemonInLocationAreaAPI(cache *pokecache.Cache, pokemonConfig *PokemonConfig, url string) error {
	data, err := handleGetRequest(cache, url)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &pokemonConfig); err != nil {
		return err
	}

	return nil
}

func GetPokemonDataAPI(cache *pokecache.Cache, pokemonDataConfig *PokemonDataConfig, url string) error {
	data, err := handleGetRequest(cache, url)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &pokemonDataConfig); err != nil {
		return err
	}

	return nil
}
