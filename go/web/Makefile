# Binary name
BINARY=passwordgen-web

# Default OS is linux
OS ?= linux

# Build directory
BUILD_DIR=build

# Go build command
GOBUILD=go build

# Version information
VERSION ?= 1.0.0
BUILD_TIME=$(shell date +%FT%T%z)

# Build flags
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

# Default target
.DEFAULT_GOAL := default

# Build for current platform
.PHONY: default
default: $(BUILD_DIR)
	$(eval XOS := $(if $(filter windows,$(OS)),windows,$(OS)))
	$(eval BINARY := $(if $(filter windows,$(OS)),passwordgen.exe,passwordgen))
	GOOS=$(XOS) $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) main.go

# Create build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Clean build directory
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Build targets for specific platforms
.PHONY: linux macos 
linux macos : $(BUILD_DIR)
	$(eval OS := $(if $(filter windows,$@),windows,$(if $(filter macos,$@),darwin,linux)))
	$(eval BINARY := $(if $(filter windows,$@),passwordgen.exe,passwordgen))
	GOOS=$(OS) $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) main.go

.PHONY: windows
windows: $(BUILD_DIR)
	GOOS=windows $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/passwordgen.exe main.go

# Build for all platforms
.PHONY: all
all: clean
	GOOS=linux $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux main.go
	GOOS=darwin $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-macos main.go
	GOOS=windows $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-windows.exe main.go

# Run the server
.PHONY: run
run: default
	./$(BUILD_DIR)/$(BINARY)

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build   - Build for current platform (default: linux)"
	@echo "  linux   - Build for Linux"
	@echo "  macos   - Build for macOS"
	@echo "  windows - Build for Windows"
	@echo "  all     - Build for all platforms"
	@echo "  clean   - Remove build directory"
	@echo "  run     - Build and run the server"
	@echo "  help    - Show this help message"
