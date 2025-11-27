package main

import (
	"testing"

	"github.com/rdawebb/pokedex-cli/internal/pokeapi"
)

func TestCatchCommandSuccess(t *testing.T) {
	cfg.pokeapiClient.Reset()
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)

	output := captureOutput(func() {
		if err := attemptCatch(cfg, "pikachu", func(max int) int { return 0 }); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	if _, exists := cfg.caughtPokemon["pikachu"]; !exists {
		t.Errorf("Expected pikachu to be caught, but it was not")
	}

	if output == "" {
		t.Errorf("Expected output, got empty string")
	}
}

func TestCatchCommandFailure(t *testing.T) {
	cfg.pokeapiClient.Reset()
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)

	output := captureOutput(func() {
		if err := attemptCatch(cfg, "bulbasaur", func(max int) int { return max }); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	if _, exists := cfg.caughtPokemon["bulbasaur"]; exists {
		t.Errorf("Expected bulbasaur to escape, but it was caught")
	}

	if output == "" {
		t.Errorf("Expected output, got empty string")
	}
}

func TestCatchCommandInvalidPokemon(t *testing.T) {
	cfg.pokeapiClient.Reset()

	err := commandCatch(cfg, []string{"invalidmon"})
	if err == nil {
		t.Errorf("Expected error for invalid Pokemon, got none")
	}
}

func TestCatchCommandNoArgs(t *testing.T) {
	cfg.pokeapiClient.Reset()

	err := commandCatch(cfg, []string{})
	if err == nil {
		t.Errorf("Expected error for no arguments, got none")
	} else {
		expectedErrMsg := "pokemon name is required"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
		}
	}
}