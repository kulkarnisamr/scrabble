package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func BenchmarkCalculateWordValue(b *testing.B) {
	// Initialize a trie
	d := &trie{root: &trieNode{children: make(map[rune]*trieNode)}}

	valuesFile, err := os.Open("value_file_2.txt")
	if err != nil {
		fmt.Println("Error opening dictionary file:", err)
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
		d.insert(key, val)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading values file:", scanner.Err())
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < 10000000; i++ {
			str := RandStringBytes(10)
			d.calculateWordValue(str)
		}
	}
}
