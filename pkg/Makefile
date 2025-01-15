.PHONY: build test clean

SUBDIRS = api util

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
	-rm -rf bin