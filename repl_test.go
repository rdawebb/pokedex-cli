package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {
            input: "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "hello world hello world",
            expected: []string{"hello", "world", "hello", "world"},
        },
        {
            input: "this is a test case",
            expected: []string{"this", "is", "a", "test", "case"},
        },
        {
            input: "",
            expected: []string{},
        },
    }
    
    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Length mismatch: %d vs %d", len(actual), len(c.expected))
            continue
        }
        
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("Word mismatch at index %d: %s vs %s", i, word, expectedWord)
            }
        }
    }
}