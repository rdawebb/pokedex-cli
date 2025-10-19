package main

import (
	"testing"
)

const expectedExploreOutput = `Exploring pastoria-city-area...
Found Pokémon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados
 - remoraid
 - octillery
 - wingull
 - pelipper
 - shellos
 - gastrodon
`

func TestExploreCommand(t *testing.T) {
	cfg.pokeapiClient.Reset()

	output := captureOutput(func() {
		if err := commandExplore(cfg, []string{"pastoria-city-area"}); err != nil {
			t.Errorf("Error executing explore command: %v", err)
		}
	})

	if output != expectedExploreOutput {
		t.Errorf("Unexpected output:\nGot:\n%s\nWant:\n%s", output, expectedExploreOutput)
	}
}

func TestExploreCommandNoArgs(t *testing.T) {
	cfg.pokeapiClient.Reset()

	err := commandExplore(cfg, []string{})
	if err == nil {
		t.Errorf("Expected error when no location area name is provided, got nil")
	} else {
		expectedErrorMsg := "location area name is required"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Unexpected error message:\nGot: %s\nWant: %s", err.Error(), expectedErrorMsg)
		}
	}
}