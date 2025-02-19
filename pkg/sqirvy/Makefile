.PHONY: debug release  test clean

debug:
	staticcheck ./...
	go vet ./...

release:
	staticcheck ./...

test:
	@echo "Testing pkg/sqirvy"
	go test .

clean:
	@echo "pkg/sqirvy"