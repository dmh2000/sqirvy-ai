.PHONY: build test clean

BINDIR := $(CURDIR)/bin

target = "sqirvy-review"

.PHONY: build test clean

BINDIR := $(CURDIR)/bin

target = "sqirvy-code"

# detect variants of Windows.PHONY: build test clean
ifeq ($(filter Windows%, $(OS)),Windows)
    arch := Windows
else
    arch := $(shell uname -s)
	ifeq ($(arch),Darwin)
		arch := darwin
	else 
		ifeq ($(arch),Linux)
			arch := linux
		else 
			echo "Unsupported OS: $(arch)"
			exit 1
		endif
	endif
endif

build: 
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=$(arch) GOARCH=amd64 go build -o $(BINDIR)/$(target) .

test:
	
clean:
	rm -rf $(BINDIR)
