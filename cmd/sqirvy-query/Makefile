.PHONY: build test clean

build:
	mkdir -p ../build
	go build -o ../build/sqirvy-query .

test:
	@./test_query.sh >/dev/null
	@./test_models.sh >/dev/null

clean:
	rm -f ../build/sqirvy-query
