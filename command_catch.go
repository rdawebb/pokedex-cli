package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func attemptCatch(cfg *config, pokemonName string, randFunc func(int) int) error {
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName, cfg.cache)
	if err != nil {
		return err
	}

	if pokemon.BaseExperience == 0 {
		return errors.New("pokemon has no base experience")
	}

	res := randFunc(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if res >= pokemon.BaseExperience/2 {
		fmt.Printf("Oh no! %s broke free!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Congratulations! You caught %s!\n", pokemon.Name)
	
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("pokemon name is required")
	}
	pokemonName := args[0]
	return attemptCatch(cfg, pokemonName, rand.Intn)
}