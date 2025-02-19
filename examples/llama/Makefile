.PHONY: debug release test clean

debug:
	staticcheck ./...
	go build -o llama .

release:
	staticcheck ./...
	mkdir -p $(BINDIR)
	go build -ldflags="-s -w" -o $(BINDIR)/llama .
	
test:
	go run . >/dev/null

clean:
	rm -f llama
	rm -f $(BINDIR)/llama
