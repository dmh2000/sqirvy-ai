.PHONY: build test clean

SUBDIRS = cmd web pkg/api

build:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir build; \
	done
	cp scripts/edit-pipe bin/edit-pipe

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
