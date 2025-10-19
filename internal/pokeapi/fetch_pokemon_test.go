package pokeapi

import (
	"testing"

	"github.com/rdawebb/pokedex-cli/internal/pokecache"
)

func TestFetchPokemonByLocationArea(t *testing.T) {
	client := NewClient()
	cache := pokecache.NewCache()
	pokemonNames, err := client.FetchPokemonByLocationArea("canalave-city-area", cache)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(pokemonNames) == 0 {
		t.Fatal("Expected to fetch some pokemon names")
	}
}