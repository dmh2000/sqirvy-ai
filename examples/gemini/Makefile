.PHONY: debug release test clean

debug:
	staticcheck ./...
	go build -o gemini .

release:
	staticcheck ./...
	mkdir -p $(BINDIR)
	go build -ldflags="-s -w" -o $(BINDIR)/gemini .

test:
	go run . >/dev/null

clean:
	rm -f gemini
	rm -f $(BINDIR)/gemini
