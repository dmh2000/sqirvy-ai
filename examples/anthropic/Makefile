.PHONY: debug release test clean

debug:
	staticcheck ./...
	go build -o anthropic .

release:
	staticcheck ./...
	mkdir -p $(BINDIR)
	go build -ldflags="-s -w" -o $(BINDIR)/anthropic .
	
test:
	go run . >/dev/null

clean:
	rm -f anthropic
	rm -f $(BINDIR)/anthropic
