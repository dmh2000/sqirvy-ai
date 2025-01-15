.PHONY: build test clean

SUBDIRS = cmd web pkg/api cmd/sqirvy-review

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
