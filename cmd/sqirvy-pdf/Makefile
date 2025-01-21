.PHONY: build test clean

# detect variants of Windows
ifeq ($(filter Windows%, $(OS)),Windows)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

build: build-linux build-macos build-windows
	@case "$$(uname -s)" in \
	Linux*) cp $(BINDIR)/sqirvy-pdf-linux $(BINDIR)/sqirvy-pdf ;; \
	Darwin*) cp $(BINDIR)/sqirvy-pdf-darwin $(BINDIR)/sqirvy-pdf ;; \
	MINGW*|MSYS*|CYGWIN*) cp $(BINDIR)/sqirvy-pdf.exe $(BINDIR)/sqirvy-pdf ;; \
	esac

build-linux:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/sqirvy-pdf-linux .

build-macos:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/sqirvy-pdf-darwin .

build-windows:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/sqirvy-pdf.exe .

test:
	@./test.sh >/dev/null
	@echo "Test Passed"

clean:
	rm -f $(BINDIR)/sqirvy-pdf-linux
	rm -f $(BINDIR)/sqirvy-pdf-darwin
	rm -f $(BINDIR)/sqirvy-pdf.exe
	rm -f $(BINDIR)/sqirvy-pdf