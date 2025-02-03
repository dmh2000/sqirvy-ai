.PHONY: build test clean

target = "sqirvy-query"
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
	@./test.sh 
	@echo "Tests passed"
	
clean:
	rm -rf $(BINDIR)
