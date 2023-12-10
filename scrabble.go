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

// dAWGNode represents a node in the dAWG
type dAWGNode struct {
	children map[rune]*dAWGNode
	value    int
	isEnd    bool
	index    int
}

// dAWG represents a directed acyclic word graph
type dAWG struct {
	root *dAWGNode
}

func newDAWG() *dAWG {
	return &dAWG{root: &dAWGNode{children: make(map[rune]*dAWGNode)}}
}

// insert inserts a word into the dAWG
func (d *dAWG) insert(word string, value int) {
	node := d.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &dAWGNode{children: make(map[rune]*dAWGNode)}
		}
		node = node.children[char]
	}
	node.value = value
	node.isEnd = true
}

// calculateWordValue calculates the value of a word by traversing the DAWG
func (d *dAWG) calculateWordValue(word string) int {
	value := 0
	i := 0
	for i < len(word) {
		found := false
		for j := len(word); j > i; j-- {
			substring := word[i:j]
			if val, ok := d.search(substring); ok {
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

// search searches for a word in the dAWG
func (d *dAWG) search(word string) (*int, bool) {
	return dfs(d.root, word)
}

// dfs performs a depth first search on the dAWG and returns true if we reach the last children nodes
func dfs(curr *dAWGNode, w string) (*int, bool) {
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

	dawg := buildDAWG(valuesFileName)

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
		value := dawg.calculateWordValue(strings.ToLower(word))
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
	fmt.Printf("total time taken to process the dawg way: %v\n", endTime.Sub(startTime).Seconds())
}

func buildDAWG(valuesFileName string) *dAWG {
	// Read values file into a dAWG
	dawg := newDAWG()
	valuesFile, err := os.Open(valuesFileName)
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
		dawg.insert(key, val)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading values file:", scanner.Err())
		os.Exit(1)
	}

	return dawg
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
