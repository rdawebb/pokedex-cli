package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type Client struct {
	baseUrl    string
	nextUrl    string
	previousUrl string
}

func NewClient() *Client {
	return &Client{
		baseUrl: "https://pokeapi.co/api/v2/location-area?limit=20",
	}
}

func (client *Client) Reset() {
		client.nextUrl = ""
		client.previousUrl = ""
}

func (client *Client) FetchLocations(url string) ([]LocationArea, error) {
	result, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch location areas: %w", err)
	}
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", result.StatusCode)
	}

	var response LocationAreaResponse
	if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	client.nextUrl = response.Next
	client.previousUrl = response.Previous

	return response.Results, nil
}

func (client *Client) GetNextUrl() string {
	return client.nextUrl
}

func (client *Client) GetPreviousUrl() string {
	return client.previousUrl
}

func (client *Client) GetBaseUrl() string {
	return client.baseUrl
}
