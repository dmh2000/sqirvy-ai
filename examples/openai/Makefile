.PHONY: build test clean

build:
	staticcheck ./...
	go build -o openai .

test:
	go run . >/dev/null

clean:
	rm -f $(BINDIR)/openai
