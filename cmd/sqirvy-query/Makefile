.PHONY: build test clean x

# detect variants of Windows
ifeq ($(filter Windows%, $(OS)),Windows)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

build: x build-linux build-macos build-windows
	@case "$$(uname -s)" in \
		Linux*) cp $(BINDIR)/sqirvy-query-linux $(BINDIR)/sqirvy-query ;; \
		Darwin*) cp $(BINDIR)/sqirvy-query-darwin $(BINDIR)/sqirvy-query ;; \
		MINGW*|MSYS*|CYGWIN*) cp $(BINDIR)/sqirvy-query.exe $(BINDIR)/sqirvy-query ;; \
	esac

build-linux:
	mkdir -p $(BINDIR)
	GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/sqirvy-query-linux .

build-macos:
	mkdir -p $(BINDIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/sqirvy-query-darwin .

build-windows:
	mkdir -p $(BINDIR)
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/sqirvy-query.exe .

test:
	@./test_query.sh >/dev/null
	@./test_models.sh >/dev/null

clean:
	rm -f $(BINDIR)/sqirvy-query-linux
	rm -f $(BINDIR)/sqirvy-query-darwin
	rm -f $(BINDIR)/sqirvy-query.exe
	rm -f $(BINDIR)/sqirvy-query
