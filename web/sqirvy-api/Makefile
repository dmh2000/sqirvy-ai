.PHONY: debug release test clean

debug:
	staticcheck ./...
	go vet ./...
	go build -o sqirvy-api .

release:
	staticcheck ./...
	go vet ./...
	mkdir -p $(BINDIR)
	go build -ldflags="-s -w" -o $(BINDIR)/sqirvy-api .

test:
	go test ./...

clean:
	rm -f sqirvy-api
	rm -f $(BINDIR)/sqirvy-api
