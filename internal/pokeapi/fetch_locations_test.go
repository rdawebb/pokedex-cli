package pokeapi

import (
	"testing"

	"github.com/rdawebb/pokedex-cli/internal/pokecache"
)

func TestClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("Expected client to be created")
	}
}

func TestFetchLocations(t *testing.T) {
	client := NewClient()
	cache := pokecache.NewCache()
	locations, err := client.FetchLocations(client.baseUrl, cache)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(locations) == 0 {
		t.Fatal("Expected to fetch some location areas")
	}
}

func TestPagination(t *testing.T) {
	client := NewClient()
	cache := pokecache.NewCache()

	// Fetch first page
	_, err := client.FetchLocations(client.baseUrl, cache)
	if err != nil {
		t.Fatalf("Expected no error on first fetch, got %v", err)
	}

	// Fetch next page
	nextUrl := client.GetNextUrl()
	if nextUrl == "" {
		t.Fatal("Expected next URL to be set after first fetch")
	}

	_, err = client.FetchLocations(nextUrl, cache)
	if err != nil {
		t.Fatalf("Expected no error on next fetch, got %v", err)
	}

	// Fetch previous page
	prevUrl := client.GetPreviousUrl()
	if prevUrl == "" {
		t.Fatal("Expected previous URL to be set after second fetch")
	}

	_, err = client.FetchLocations(prevUrl, cache)
	if err != nil {
		t.Fatalf("Expected no error on previous fetch, got %v", err)
	}
}