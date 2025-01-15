.PHONY: build test clean

build:
	staticcheck ./...
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/openai .

test:
	echo "Skipping tests for openai"

clean:
	rm -f $(BINDIR)/openai
