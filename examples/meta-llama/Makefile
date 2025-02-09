.PHONY: build test clean

build:
	staticcheck ./...
	go build -o $(BINDIR)/meta-llama .
	
test:
	go run . >/dev/null

clean:
	rm -f $(BINDIR)/meta-llama
