.PHONY: build run clean

build:
	mkdir -p build
	go build -o build/sqirvy-web main.go
	cp -r static build/

run: build
	./build/server

test:
	echo "Testinng examples/web"

clean:
	rm -rf build
