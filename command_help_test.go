package main

import (
	"strings"
	"testing"
)

func TestHelpCommand(t *testing.T) {
		output := captureOutput(func() {
			if err := commandHelp(); err != nil {
				t.Errorf("Error executing help command: %v", err)
			}
		})
		
		if !strings.Contains(output, "Welcome to the Pokedex!") {
			t.Errorf("Help output does not contain welcome message")
		}

		if !strings.Contains(output, "Usage:") {
			t.Errorf("Help output does not contain usage message")
		}

		expectedCommands := []string{
			"help: Displays a help message",
			"map: Displays a list of 20 location areas",
			"mapb: Displays a list of the previous 20 location areas",
			"exit: Exit the Pokedex",
		}

		for _, cmd := range expectedCommands {
			if !strings.Contains(output, cmd) {
				t.Errorf("Help output does not contain command: %s", cmd)
			}
		}

		if !strings.HasPrefix(output, "\n") || !strings.HasSuffix(output, "\n") {
			t.Errorf("Help output does not have expected leading and trailing newlines")
		}
	}