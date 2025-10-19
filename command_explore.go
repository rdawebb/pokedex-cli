package main

import (
	"fmt"
)

func fetchPokemon(location string) error {
	pokemon, err := cfg.pokeapiClient.FetchPokemonByLocationArea(location, cfg.cache)
	if err != nil {
		return err
	}

	if len(pokemon) == 0 {
		fmt.Println("No Pokémon found in this location area.")
		return nil
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokémon:")

	for _, pokemonName := range pokemon {
		fmt.Println(" - " + pokemonName)
	}

	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("location area name is required")
	}
	location := args[0]
	return fetchPokemon(location)
}