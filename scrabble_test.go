package main

import (
	"testing"
)

func TestWordValue(t *testing.T) {
	// Initialize a trie
	trie, err := buildTrie("value_file_2.txt")
	if err != nil {
		panic(err)
	}
	// Define test cases
	tests := []struct {
		name  string
		word  string
		value int
	}{
		{"Value of hacker", "bushwhacker", 22},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trie.calculateWordValue(tt.word); got != tt.value {
				t.Errorf("calculateWordValue() = %v, want %v", got, tt.value)
			}
		})
	}
}
