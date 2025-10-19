package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rdawebb/pokedex-cli/internal/pokeapi"
	"github.com/rdawebb/pokedex-cli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient *pokeapi.Client
	cache         *pokecache.Cache
}

var cfg *config

func init() {
	cfg = &config{}
	cfg.pokeapiClient = pokeapi.NewClient()
	cfg.cache = pokecache.NewCache()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a list of the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location area",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			line := scanner.Text()
			words := cleanInput(line)

			if len(words) == 0 {
				continue
			}

			command := words[0]
			args := words[1:]

			if cmd, exists := getCommands()[command]; exists {
				if err := cmd.callback(cfg, args); err != nil {
					fmt.Println("Error executing command:", err)
				}
			} else {
				fmt.Println("Unknown command:", command)
			}

		} else {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			break
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
    return words
}
