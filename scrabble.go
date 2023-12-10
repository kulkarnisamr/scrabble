package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// trieNode represents a node in the trie
type trieNode struct {
	children map[rune]*trieNode
	value    int
	isEnd    bool
}

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{root: &trieNode{children: make(map[rune]*trieNode)}}
}

// insert inserts a word into the trie
func (t *trie) insert(word string, value int) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &trieNode{children: make(map[rune]*trieNode)}
		}
		node = node.children[char]
	}
	node.value = value
	node.isEnd = true
}

// calculateWordValue calculates the value of a word by traversing the DAWG
func (t *trie) calculateWordValue(word string) int {
	value := 0
	i := 0
	for i < len(word) {
		found := false
		for j := len(word); j > i; j-- {
			substring := word[i:j]
			if val, ok := t.search(substring); ok {
				value += *val
				i = j
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}
	return value
}

// search searches for a word in the trie
func (t *trie) search(word string) (*int, bool) {
	return dfs(t.root, word)
}

// dfs performs a depth first search on the trie and returns true if we reach the last children nodes
func dfs(curr *trieNode, w string) (*int, bool) {
	for i := range w {
		if _, ok := curr.children[rune(w[i])]; !ok {
			return nil, false
		}

		curr = curr.children[rune(w[i])]
	}

	return &curr.value, curr.isEnd
}

// node represents a word and its value (only used as a helper for printing)
type node struct {
	key   string
	value int
}

var valuesFile = flag.String("values", "", "input values file")
var dictFile = flag.String("dict", "", "input dict file")

func main() {
	flag.Parse()
	startTime := time.Now().UTC()
	if len(os.Args) != 3 {
		fmt.Println("Usage: program_name dictionary_file value_file")
		os.Exit(1)
	}

	var dictionaryFileName, valuesFileName string
	if *dictFile != "" {
		dictionaryFileName = *dictFile
	} else {
		log.Fatal("dict file not specified")
	}

	if *valuesFile != "" {
		valuesFileName = *valuesFile
	} else {
		log.Fatal("values file not specified")
	}

	trie, err := buildTrie(valuesFileName)
	if err != nil {
		panic("failed to build a trie")
	}

	// read dict file
	dict, err := os.Open(dictionaryFileName)
	if err != nil {
		panic("failed to open dict file")
	}
	defer dict.Close()

	result := make(map[string]node)
	highestValue := 0

	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		word := scanner.Text()
		value := trie.calculateWordValue(strings.ToLower(word))
		if value > highestValue {
			result[word] = node{
				key:   word,
				value: value,
			}
			highestValue = value
		} else {
			if value == highestValue {
				result[word] = node{
					key:   word,
					value: value,
				}
			}
		}
	}

	printHighestValueWords(result, highestValue)

	endTime := time.Now().UTC()
	fmt.Printf("total time taken to process the trie way: %v\n", endTime.Sub(startTime).Seconds())
}

func buildTrie(valuesFileName string) (*trie, error) {
	// Read values file into a trie
	t := newTrie()
	valuesFile, err := os.Open(valuesFileName)
	if err != nil {
		fmt.Println("Error opening values file:", err)
		return nil, err
	}
	defer valuesFile.Close()

	scanner := bufio.NewScanner(valuesFile)
	for scanner.Scan() {
		var key string
		var val int
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &key, &val)
		if err != nil {
			fmt.Println("Error parsing values file:", err)
			return nil, err
		}
		t.insert(strings.ToLower(key), val)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading values file:", scanner.Err())
		return nil, err
	}

	return t, nil
}

func printHighestValueWords(result map[string]node, highestValue int) {
	resultArr := make([]node, 0, len(result))
	for _, node := range result {
		resultArr = append(resultArr, node)
	}

	sort.Slice(resultArr, func(i, j int) bool {
		return resultArr[i].value > resultArr[j].value
	})

	for _, res := range resultArr {
		if res.value == highestValue {
			fmt.Printf("%s %d\n", res.key, res.value)
		} else {
			break
		}
	}
}
