.PHONY: build test clean

BINDIR := $(CURDIR)/bin

target = "sqirvy-query"

# detect variants of Windows
ifeq ($(filter Windows%, $(OS)),Windows)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

ifeq ($(DETECTED_OS),Windows)
	BUILD_TARGET := build-windows
else ifeq ($(DETECTED_OS),Darwin)
	BUILD_TARGET := build-macos
else ifeq ($(DETECTED_OS),Linux)
	BUILD_TARGET := build-linux
else
	$(error "Unsupported OS: $(DETECTED_OS)")
endif

build: $(BUILD_TARGET)
	@echo "Build target: $(BUILD_TARGET)"

build-linux:
	staticcheck ./...
	mkdir -p $(BINDIR)
	echo "Building for Linux" 
	GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/$(target) . 

build-macos:
	staticcheck ./...
	mkdir -p $(BINDIR)
	echo "Building for MacOS"
	GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/$(target) . 

build-windows:
	staticcheck ./...
	mkdir -p $(BINDIR)
	echo "Building for Windows"; 
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/$(target) . 

test:
	@./test_models.sh >/dev/null
	@./test_query.sh >/dev/null

clean:
	rm -rf $(BINDIR)

