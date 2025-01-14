.PHONY: build test clean

build: build-linux build-macos build-windows
	@case "$$(uname -s)" in \
		Linux*) cp ../build/sqirvy-query-linux ../build/sqirvy-query ;; \
		Darwin*) cp ../build/sqirvy-query-darwin ../build/sqirvy-query ;; \
		MINGW*|MSYS*|CYGWIN*) cp ../build/sqirvy-query.exe ../build/sqirvy-query ;; \
	esac

build-linux:
	mkdir -p ../build
	GOOS=linux GOARCH=amd64 go build -o ../build/sqirvy-query-linux .

build-macos:
	mkdir -p ../build
	GOOS=darwin GOARCH=amd64 go build -o ../build/sqirvy-query-darwin .

build-windows:
	mkdir -p ../build
	GOOS=windows GOARCH=amd64 go build -o ../build/sqirvy-query.exe .

test:
	@./test_query.sh >/dev/null
	@./test_models.sh >/dev/null

clean:
	rm -f ../build/sqirvy-query-linux
	rm -f ../build/sqirvy-query-darwin
	rm -f ../build/sqirvy-query.exe
	rm -f ../build/sqirvy-query
