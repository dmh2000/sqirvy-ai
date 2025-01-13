.PHONY: build test clean

SUBDIRS = cmd web pkg/api

build:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir build; \
	done

test:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir test; \
	done

clean:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir clean; \
	done
