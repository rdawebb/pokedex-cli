package main

import (
	"testing"
)

const expectedMapOutput = `canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
`

func TestMapCommand(t *testing.T) {
	cfg.pokeapiClient.Reset()

	output := captureOutput(func() {
		if err := commandMap(); err != nil {
			t.Errorf("Error executing map command: %v", err)
		}
	})
	
	if output != expectedMapOutput {
		t.Errorf("Unexpected output:\nGot:\n%s\nWant:\n%s", output, expectedMapOutput)
	}
}

func TestMapbCommand(t *testing.T) {
	cfg.pokeapiClient.Reset()

	// Call commandMap twice to get to the second page
	if err := commandMap(); err != nil {
		t.Errorf("Error executing map command: %v", err)
	}

	if err := commandMap(); err != nil {
		t.Errorf("Error executing map command: %v", err)
	}

	output := captureOutput(func() {
		if err := commandMapb(); err != nil {
			t.Errorf("Error executing mapb command: %v", err)
		}
	})

	if output != expectedMapOutput {
		t.Errorf("Unexpected output:\nGot:\n%s\nWant:\n%s", output, expectedMapOutput)
	}
}

func TestMapbCommandFirstPage(t *testing.T) {
	cfg.pokeapiClient.Reset()

	output := captureOutput(func() {
		if err := commandMapb(); err != nil {
			t.Errorf("Error executing mapb command: %v", err)
		}
	})

	expectedOutput := "You're on the first page.\n"
	if output != expectedOutput {
		t.Errorf("Unexpected output:\nGot:\n%s\nWant:\n%s", output, expectedOutput)
	}
}
