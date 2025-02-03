.PHONY: build test clean

build:
	staticcheck ./...
	go build -o $(BINDIR)/gemini .

test:
	go run . >/dev/null

clean:
	rm -f $(BINDIR)/gemini
