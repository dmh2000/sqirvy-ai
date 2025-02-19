.PHONY: debug release test clean

debug:
	staticcheck ./...
	go build -o openai .

release:
	staticcheck ./...
	mkdir -p $(BINDIR)
	go build -ldflags="-s -w" -o $(BINDIR)/openai .

test:
	go run . >/dev/null

clean:
	rm -f openai
	rm -f $(BINDIR)/openai
