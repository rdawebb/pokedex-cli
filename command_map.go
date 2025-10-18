package main

import (
	"fmt"
)

func fetchLocations(url string) error {
	areas, err := cfg.pokeapiClient.FetchLocations(url)
	if err != nil {
		return err
	}

	if len(areas) == 0 {
		fmt.Println("No location areas found.")
		return nil
	}

	for _, area := range areas {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMap() error {
	url := cfg.pokeapiClient.GetNextUrl()
	if url == "" {
		url = cfg.pokeapiClient.GetBaseUrl()
	}
	return fetchLocations(url)
}

func commandMapb() error {
	previousUrl := cfg.pokeapiClient.GetPreviousUrl()
	if previousUrl == "" {
		fmt.Println("You're on the first page.")
		return nil
	}
	return fetchLocations(previousUrl)
}