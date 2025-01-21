.PHONY: build test clean

# detect variants of Windows
ifeq ($(filter Windows%, $(OS)),Windows)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

build: build-linux build-macos build-windows
	@case "$$(uname -s)" in \
		Linux*) cp $(BINDIR)/sqirvy-query-linux $(BINDIR)/sqirvy-query ;; \
		Darwin*) cp $(BINDIR)/sqirvy-query-darwin $(BINDIR)/sqirvy-query ;; \
		MINGW*|MSYS*|CYGWIN*) cp $(BINDIR)/sqirvy-query.exe $(BINDIR)/sqirvy-query ;; \
	esac

build-linux:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/sqirvy-query-linux .

build-macos:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/sqirvy-query-darwin .

build-windows:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/sqirvy-query.exe .

test:
	@./test_query.sh >/dev/null
	@./test_models.sh >/dev/null
	@echo "Test Passed"

clean:
	rm -f $(BINDIR)/sqirvy-query-linux
	rm -f $(BINDIR)/sqirvy-query-darwin
	rm -f $(BINDIR)/sqirvy-query.exe
	rm -f $(BINDIR)/sqirvy-query
