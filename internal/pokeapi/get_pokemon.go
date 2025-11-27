package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rdawebb/pokedex-cli/internal/pokecache"
)

func (client *Client) GetPokemon(pokemonName string, cache *pokecache.Cache) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	if data, ok := cache.Get(url); ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(data, &pokemonResponse)
		if err == nil {
			return pokemonResponse, nil
		}
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Pokemon{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResponse := Pokemon{}
	err = json.Unmarshal(body, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}

	cache.Add(url, body)

	return pokemonResponse, nil
}
	
