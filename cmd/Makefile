.PHONY: build test clean

SUBDIRS = sqirvy-query sqirvy-review sqirvy-scrape gemini openai anthropic
export BINDIR:=../../bin

build:
	@for dir in $(SUBDIRS); do \
		$(MAKE) $(SILENT)  -C $$dir build; \
	done

test:
	@for dir in $(SUBDIRS); do \
		echo "Testing $$dir"; \
		$(MAKE) $(SILENT)  -C $$dir test; \
	done

clean:
	@for dir in $(SUBDIRS); do \
		$(MAKE) $(SILENT)  -C $$dir clean; \
	done
	rm -rf build
