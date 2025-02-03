.PHONY: build test clean

build:
	staticcheck ./...
	go build -o $(BINDIR)/openai .

test:
	go run . >/dev/null

clean:
	rm -f $(BINDIR)/openai
