.PHONY: debug release test clean

debug:
	staticcheck ./...
	go vet ./...
	go build -o sqirvy-xyz .

release:
	staticcheck ./...
	go vet ./...
	mkdir -p $(BINDIR)
	go build -ldflags="-s -w" -o $(BINDIR)/sqirvy-xyz .

test:
	@echo "no tests in sqirvy-xyz"

clean:
	rm -f sqirvy-xyz
	rm -f $(BINDIR)/sqirvy-xyz
