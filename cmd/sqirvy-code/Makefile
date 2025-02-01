.PHONY: build test clean
# detect variants of Windows
ifeq ($(filter Windows%, $(OS)),Windows)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

build: build-linux build-macos build-windows
	@case "$$(uname -s)" in \
		Linux*) cp $(BINDIR)/sqirvy-code-linux $(BINDIR)/sqirvy-code ;; \
		Darwin*) cp $(BINDIR)/sqirvy-code-darwin $(BINDIR)/sqirvy-code ;; \
		MINGW*|MSYS*|CYGWIN*) cp $(BINDIR)/sqirvy-code.exe $(BINDIR)/sqirvy-code ;; \
	esac

build-linux:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/sqirvy-code-linux .

build-macos:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/sqirvy-code-darwin .

build-windows:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/sqirvy-code.exe .

test:
	@echo "skip tests for sqirvy-code"

clean:
	rm -f $(BINDIR)/sqirvy-code-linux
	rm -f $(BINDIR)/sqirvy-code-darwin
	rm -f $(BINDIR)/sqirvy-code.exe
	rm -f $(BINDIR)/sqirvy-code
