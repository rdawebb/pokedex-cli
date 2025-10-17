package main

import (
	"fmt"
    "strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
    return words
}