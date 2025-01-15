.PHONY: build test clean

build:
	staticcheck ./...
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/anthropic .
	
test:
	echo "Skipping tests for anthropic"

clean:
	rm -f $(BINDIR)/anthropic
