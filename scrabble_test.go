package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestWordValue(t *testing.T) {
	// Initialize a trie
	t1 := &trie{root: &trieNode{children: make(map[rune]*trieNode)}}

	valuesFile, err := os.Open("value_file_2.txt")
	if err != nil {
		fmt.Println("Error opening values file:", err)
		os.Exit(1)
	}
	defer valuesFile.Close()

	scanner := bufio.NewScanner(valuesFile)
	for scanner.Scan() {
		var key string
		var val int
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &key, &val)
		if err != nil {
			fmt.Println("Error parsing values file:", err)
			os.Exit(1)
		}
		t1.insert(key, val)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading values file:", scanner.Err())
		os.Exit(1)
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
			if got := t1.calculateWordValue(tt.word); got != tt.value {
				t.Errorf("calculateWordValue() = %v, want %v", got, tt.value)
			}
		})
	}
}
