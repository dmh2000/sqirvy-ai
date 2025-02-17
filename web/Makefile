.PHONY: build test clean

SUBDIRS = sqirvy-api sqirvy-xyz

build:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir build; \
	done

test:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir test; \
	done

clean:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir clean; \
	done
	-rm -rf bin