.PHONY: build test clean
# detect variants of Windows
ifeq ($(filter Windows%, $(OS)),Windows)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

build: build-linux build-macos build-windows
	@case "$$(uname -s)" in \
		Linux*) cp $(BINDIR)/sqirvy-review-linux $(BINDIR)/sqirvy-review ;; \
		Darwin*) cp $(BINDIR)/sqirvy-review-darwin $(BINDIR)/sqirvy-review ;; \
		MINGW*|MSYS*|CYGWIN*) cp $(BINDIR)/sqirvy-review.exe $(BINDIR)/sqirvy-review ;; \
	esac

build-linux:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/sqirvy-review-linux .

build-macos:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/sqirvy-review-darwin .

build-windows:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/sqirvy-review.exe .

test:
	@echo "skip tests for sqirvy-review"

clean:
	rm -f $(BINDIR)/sqirvy-review-linux
	rm -f $(BINDIR)/sqirvy-review-darwin
	rm -f $(BINDIR)/sqirvy-review.exe
	rm -f $(BINDIR)/sqirvy-review
