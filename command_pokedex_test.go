package main

import (
	"strings"
	"testing"

	"github.com/rdawebb/pokedex-cli/internal/pokeapi"
)

func TestCommandPokedex(t *testing.T) {
	cfg.pokeapiClient.Reset()
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)

	// Simulate catching some Pokemon
	cfg.caughtPokemon["pikachu"] = pokeapi.Pokemon{Name: "pikachu"}
	cfg.caughtPokemon["bulbasaur"] = pokeapi.Pokemon{Name: "bulbasaur"}

	output := captureOutput(func() {
		if err := commandPokedex(cfg, []string{}); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	if !strings.Contains(output, "Your Pokedex:") {
		t.Error("Expected 'Your Pokedex:' header in output")
	}

	expectedOutputs := []string{"pikachu", "bulbasaur"}
	for _, expected := range expectedOutputs {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain %s, but it did not", expected)
		}
	}
}

func TestCommandPokedexEmpty(t *testing.T) {
	cfg.pokeapiClient.Reset()
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)

	output := captureOutput(func() {
		if err := commandPokedex(cfg, []string{}); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	if !strings.Contains(output, "Your Pokedex:") {
		t.Error("Expected 'Your Pokedex:' header in output")
	}
}