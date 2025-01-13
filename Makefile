.PHONY: build test clean

SUBDIRS = cmd web pkg/api

build:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir build; \
	done

test:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir test; \
	done
	@echo "Tests passed"

clean:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir clean; \
	done
