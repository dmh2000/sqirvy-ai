.PHONY: build run clean

build:
	mkdir -p build
	go build -o build/sqirvy-web main.go
	cp -r static build/

run: build
	./build/server

test:
	echo "Skipping tests for web"

clean:
	rm -rf build
