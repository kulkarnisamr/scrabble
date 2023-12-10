GO=go

tidy:
	$(GO) mod tidy -compat=1.20

cover:
	$(GO) test ./...

build:
	$(GO) build scrabble.go

collins:
	$(GO) run scrabble.go -dict=collins.txt -values=scrabble_tile_values.txt

benchmark:
	$(GO) test -bench=. -benchmem ./...