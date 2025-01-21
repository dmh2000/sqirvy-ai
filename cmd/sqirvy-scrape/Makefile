# detect variants of Windows
ifeq ($(filter Windows%, $(OS)),Windows)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

build: build-linux build-macos build-windows
	@case "$$(uname -s)" in \
		Linux*) cp $(BINDIR)/sqirvy-scrape-linux $(BINDIR)/sqirvy-scrape ;; \
		Darwin*) cp $(BINDIR)/sqirvy-scrape-darwin $(BINDIR)/sqirvy-scrape ;; \
		MINGW*|MSYS*|CYGWIN*) cp $(BINDIR)/sqirvy-scrape.exe $(BINDIR)/sqirvy-scrape ;; \
	esac

build-linux:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/sqirvy-scrape-linux .

build-macos:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/sqirvy-scrape-darwin .

build-windows:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/sqirvy-scrape.exe .

test:
	@./test.sh >/dev/null
	@echo "Test Passed"


clean:
	rm -f $(BINDIR)/sqirvy-scrape-linux
	rm -f $(BINDIR)/sqirvy-scrape-darwin
	rm -f $(BINDIR)/sqirvy-scrape.exe
	rm -f $(BINDIR)/sqirvy-scrape
