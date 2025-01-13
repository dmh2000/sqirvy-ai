.PHONY: build test clean

SUBDIRS = sqirvy-query gemini openai anthropic

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
