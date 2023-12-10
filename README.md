# scrabble
A version of scrabble that uses Trie and prints the highest valued words given a dictionary of words and a file with single characters or strings and their tile values

Note: The final list of words is not lexicographically sorted

### How to run
```
make build
./scrabble -dict=dictionary.txt -values=value_file_2.txt
```
OR
```
go run ./scrabble.go -dict=dictionary.txt -values=value_file_2.txt
```

You can also run a default version of the program which uses Collins dictionary words
and original scrabble tile values by running:
```
make collins
```

### Unit tests
```
make cover
```

### Benchmarks
```
make benchmark
```
On my machine for 10M random words:
```
go test -bench=. -benchmem ./...
goos: darwin
goarch: amd64
pkg: github.com/kulkarnisamr/scrabble
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkCalculateWordValue-16    	       1	14060423618 ns/op	320016928 B/op	20000111 allocs/op
PASS
ok  	github.com/kulkarnisamr/scrabble	14.232s
```

# Performance
### Go version
`go 1.20.5`

### Runtime
```
Hardware Overview:

Model Name:	MacBook Pro
Model Identifier:	MacBookPro16,1
Processor Name:	8-Core Intel Core i9
Processor Speed:	2.4 GHz
Number of Processors:	1
Total Number of Cores:	8
L2 Cache (per Core):	256 KB
L3 Cache:	16 MB
Hyper-Threading Technology:	Enabled
Memory:	32 GB
```

### Complexity
1. **Insertion**: The insertion operation involves traversing the trie to insert a new word. The time complexity is O(m), where m is the length of the word being inserted.
2. **CalculateWordValue**: The time complexity of calculating the value of a word is O(n), where n is the length of the word. This is because we traverse the Trie from the root to the corresponding node for the word.
3. **Space Complexity**: The space complexity of a trie is influenced by the total number of nodes and the average length of an entry in the values file. For a trie with n keys, the space complexity is O(n * m), where m is the average length of the keys.

Overall, the time complexity of constructing the Trie (insertion) is influenced by the total number of vertices and edges in the trie. 
The time complexity for looking up the value of a word is primarily influenced by the length of the word.

### Approach:
**Trie**: Trie is efficient for string search operations where prefixes are shared and we always start at the beginning of the word.

### Options considered:
**DAWG**: Directed Acyclic Word Graph is a data structure that is used to store a set of strings. It is a compacted trie. With the current input files,
there isn't a clear advantage between using a DAWG or a Trie so I decided to avoid it because of the complexity of the compaction operation. However, if the input file were large, DAWG will be a better option as it allows suffix sharing and the number of nodes and edges reduce significantly.

DAWG Benchmark for 10M random words:
```
go test -bench=. -benchmem ./...
goos: darwin
goarch: amd64
pkg: github.com/kulkarnisamr/scrabble
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkCalculateWordValue-16    	       1	14366705763 ns/op	320017848 B/op	20000117 allocs/op
PASS
ok  	github.com/kulkarnisamr/scrabble	14.538s
```

### Pros
1. **Efficient Prefix Searches**: Tries are excellent for prefix searches as they can quickly identify all strings with a given prefix.
2. **Fast insertion and lookup**: Insertion and search operations in a trie have a time complexity of O(m), where m is the length of the string. This makes trie operations efficient for string-related tasks.
3. **Keys are ordered**: Tries maintain a lexicographical order of the keys, making it easy to iterate over all keys in sorted order.
4. **Simple implementation**: Implementation is quite easy compared to a DAWG

### Cons
1. **Space efficiency**: If the input value files were large, DAWGs would be more space-efficient compared to plain tries because they allow suffix sharing.
2. **High Memory Consumption**: Tries can have higher memory consumption compared to DAWGs, especially when the input set is large or when there are many keys with distinct prefixes.

### Further improvements:
Add DAWG implementation which will bring forth its advantages with space compaction when we use large files.

### References:
https://www.cs.cmu.edu/afs/cs/academic/class/15451-s06/www/lectures/scrabble.pdf

