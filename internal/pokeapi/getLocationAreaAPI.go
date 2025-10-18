package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/delroscol98/pokedex/internal/pokecache"
)

func GetLocationAreaAPI(cache *pokecache.Cache, config *LocationConfig, url string) error {
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
			return err
		}

		data, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return fmt.Errorf("error reading response: %w", err)
		}
	}

	cache.Add(url, data)

	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	return nil
}
