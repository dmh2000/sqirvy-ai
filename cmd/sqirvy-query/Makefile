.PHONY: build test clean x

# detect variants of Windows
ifeq ($(filter Windows%, $(OS)),Windows)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

build: x build-linux build-macos build-windows
	@case "$$(uname -s)" in \
		Linux*) cp ../../bin/sqirvy-query-linux ../../bin/sqirvy-query ;; \
		Darwin*) cp ../../bin/sqirvy-query-darwin ../../bin/sqirvy-query ;; \
		MINGW*|MSYS*|CYGWIN*) cp ../../bin/sqirvy-query.exe ../../bin/sqirvy-query ;; \
	esac

build-linux:
	mkdir -p ../../bin
	GOOS=linux GOARCH=amd64 go build -o ../../bin/sqirvy-query-linux .

build-macos:
	mkdir -p ../../bin
	GOOS=darwin GOARCH=amd64 go build -o ../../bin/sqirvy-query-darwin .

build-windows:
	mkdir -p ../../bin
	GOOS=windows GOARCH=amd64 go build -o ../../bin/sqirvy-query.exe .

test:
	@./test_query.sh >/dev/null
	@./test_models.sh >/dev/null

clean:
	rm -f ../../bin/sqirvy-query-linux
	rm -f ../../bin/sqirvy-query-darwin
	rm -f ../../bin/sqirvy-query.exe
	rm -f ../../bin/sqirvy-query
