.PHONY: debug release test clean

target = "sqirvy-query"

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
	GOOS=$(arch) GOARCH=amd64 go build -o $(target) .

release:
	staticcheck ./...
	mkdir -p $(BINDIR)
	GOOS=$(arch) GOARCH=amd64 go build -ldflags="-s -w" -o $(BINDIR)/$(target) .

test:
	@./test.sh 
	@echo "Tests passed"
	
clean:
	rm -f $(target)
	rm -rf $(BINDIR)
