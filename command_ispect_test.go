package main

import (
	"testing"

	"github.com/rdawebb/pokedex-cli/internal/pokeapi"
)

func TestInspectCommandSuccess(t *testing.T) {
	cfg.pokeapiClient.Reset()
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)

	if err := attemptCatch(cfg, "pikachu", func(max int) int { return 0 }); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err := commandInspect(cfg, []string{"pikachu"})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestInspectCommandFailure(t *testing.T) {
	cfg.pokeapiClient.Reset()
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)

	err := commandInspect(cfg, []string{"bulbasaur"})
	if err == nil {
		t.Errorf("Expected error for un-caught Pokemon, got none")
	}
}

func TestInspectCommandNoArgs(t *testing.T) {
	cfg.pokeapiClient.Reset()

	err := commandInspect(cfg, []string{})
	if err == nil {
		t.Errorf("Expected error for no arguments, got none")
	} else {
		expectedErrMsg := "pokemon name is required"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
		}
	}
}