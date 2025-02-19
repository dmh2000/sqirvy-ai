.PHONY: debug release test clean

debug:
	staticcheck ./...
	go vet ./...


release:
	staticcheck ./...
	go vet ./...


test:
	@echo "Testing pkg/util"
	go test .

clean:
	@echo "pkg/util"