.PHONY: build test clean

build:
	mkdir -p ../build
	go build -o ../build/sqirvy-query .

test:
	@./test.sh >/dev/null

clean:
	rm -f ../build/sqirvy-query
