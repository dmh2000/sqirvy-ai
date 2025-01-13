.PHONY: build run clean

build:
	mkdir -p build
	go build -o build/server main.go
	cp -r static build/

run: build
	./build/server

clean:
	rm -rf build
