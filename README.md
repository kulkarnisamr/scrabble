# scrabble
A version of scrabble that prints the highest valued words given a dictionary of words and a file with single characters or strings and their tile values

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
BenchmarkCalculateWordValue-16    	       1	14366705763 ns/op	320017848 B/op	20000117 allocs/op
PASS
ok  	github.com/kulkarnisamr/scrabble	14.538s
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
2. **CalculateWordValue**: The time complexity of calculating the value of a word is O(n), where n is the length of the word. This is because we traverse the DAWG from the root to the corresponding node for the word.

Overall, the time complexity of constructing the DAWG (insertion and compaction) is influenced by the total number of vertices and edges in the DAWG. 
The time complexity for looking up the value of a word is primarily influenced by the length of the word.

### Approach:
**DAWG**: Directed Acyclic Word Graph is a data structure that is used to store a set of strings. It is a compacted trie. With the current input files,
there isn't a clear advantage between using a DAWG or a Trie so the current implementation is just a glorified Trie since compaction is not implemented. However, if the input file were large, the DAWG will be more compact and hence will take less space making it easier to extend the current functionality.

### Options considered:
1. **Trie**: The implementation is very similar and the time complexity is the same as DAWG. The benchmarks were also comparable.
```
goos: darwin
goarch: amd64
pkg: github.com/kulkarnisamr/single_reader/trie
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkCalculateWordValue-16    	       1	14146373726 ns/op	320015224 B/op	20000109 allocs/op
PASS
ok  	github.com/kulkarnisamr/single_reader/trie	14.411s

```
### Pros
1. **Space efficiency**: If the input value files were large, DAWGs would be more space-efficient compared to plain tries because they allow suffix sharing.
2. **Reduced Memory Usage**: Since suffixes are shared, DAWG's reduce the total number of nodes and edges significantly. For large dictionaries this can make a huge difference.
3. **Improved Lookup Performance**: DAWGs can lead to faster lookup times for string search operations as they eliminate redundancy and allow for more efficient traversal.

### Cons
1. Implementation complexity: The implementation of DAWG is more complex than a trie because of the compaction step (not added in this implementation).
2. The additional compaction step is also computationally expensive as we need to reindex nodes.

### Further improvements:
Add compaction method to the DAWG implementation which will bring forth its advantages when we use large files.

### References:
https://www.cs.cmu.edu/afs/cs/academic/class/15451-s06/www/lectures/scrabble.pdf

