package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rdawebb/pokedex-cli/internal/pokecache"
)

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
}

type PokemonEncountersResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func (client *Client) FetchPokemonByLocationArea(location string, cache *pokecache.Cache) ([]string, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	if data, ok := cache.Get(url); ok {
		var response PokemonEncountersResponse
		if err := json.Unmarshal(data, &response); err == nil {
			var pokemonNames []string
			for _, encounter := range response.PokemonEncounters {
				pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
			}
			return pokemonNames, nil
		}
	}

	result, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pokemon encounters: %w", err)
	}
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", result.StatusCode)
	}

	var response PokemonEncountersResponse
	if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var pokemonNames []string
	for _, encounter := range response.PokemonEncounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}

	if data, err := json.Marshal(response); err == nil {
		cache.Add(url, data)
	}

	return pokemonNames, nil
}
