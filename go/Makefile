# Default target
.DEFAULT_GOAL := default

.PHONY: default
default: cli web

# Build CLI application
.PHONY: cli
cli:
	cd cmd/cli && $(MAKE)

# Build web application
.PHONY: web
web:
	cd web && $(MAKE)

# Build both CLI and web applications
.PHONY: all
all: cli web
	cd cmd/cli && $(MAKE) all
	cd web && $(MAKE) all

# Clean both applications
.PHONY: clean
clean:
	cd cmd/cli && $(MAKE) clean
	cd web && $(MAKE) clean

# Run the web server
.PHONY: run-web
run-web:
	cd web && $(MAKE) run

# Run CLI help
.PHONY: run-cli
run-cli:
	cd cmd/cli && $(MAKE) && ./build/passwordgen -h

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  ''       - Build both CLI and web applications for current platform (default)"
	@echo "  all      - Build both CLI and web applications (default)"
	@echo "  cli      - Build CLI application"
	@echo "  web      - Build web application"
	@echo "  clean    - Clean both applications"
	@echo "  run-web  - Build and run the web server"
	@echo "  run-cli  - Build and run the CLI (shows help)"
	@echo "  help     - Show this help message"
