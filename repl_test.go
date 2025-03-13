package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "CHARmander BulbaSAUR PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	passed := 0
	failed := 0

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Test Failed: expected %d words, got %d", len(c.expected), len(actual))
			failed++
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Test Failed: expected %s, got %s", expectedWord, word)
				failed++
			}
			fmt.Println("Test Passed!")
			passed++
		}

		fmt.Println("---------------------------------")
		fmt.Printf("%d passed, %d failed\n", passed, failed)
	}
}
