.PHONY: build test clean

build:
	staticcheck ./...
	go build -o $(BINDIR)/llama .
	
test:
	go run . >/dev/null

clean:
	rm -f $(BINDIR)/llama
