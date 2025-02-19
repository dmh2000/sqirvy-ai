.PHONY: debug release test clean

PROJECT=sqirvy-cli
BINDIR=../../bin

# detect variants of Windows
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

debug: 
	staticcheck ./...
	GOOS=$(arch) GOARCH=amd64 go build -o $(PROJECT) .

release: 
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=$(arch) GOARCH=amd64 go build -licflags="-s -w" -o $(BINDIR)/$(PROJECT) .

test:
	@echo "sqirvy-cli no tests"
	
clean:
	rm -rf $(PROJECT)
	rm -rf $(BINDIR)
	rm -rf ./tmp
