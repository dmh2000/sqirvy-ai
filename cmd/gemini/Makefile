.PHONY: build test clean

build:
	staticcheck ./...
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/gemini .

test:
	echo "Skipping tests for gemini"

clean:
	rm -f $(BINDIR)/gemini
