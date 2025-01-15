.PHONY: build test clean

build:
	staticcheck ./...
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/anthropic .
	
test:
	go run . >/dev/null

clean:
	rm -f $(BINDIR)/anthropic
