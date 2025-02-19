.PHONY: debug release test clean

SUBDIRS = sqirvy-review sqirvy-code sqirvy-scrape anthropic gemini openai

export BINDIR=$(CURDIR)/build

debug:
	@for dir in $(SUBDIRS); do \
		$(MAKE) $(SILENT)  -C $$dir debug; \
	done

release:
	@echo "no release builds"

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