.PHONY: debug release test clean

SUBDIRS = sqirvy-api sqirvy-xyz

export BINDIR=$(CURDIR)/build

debug:
	@for dir in $(SUBDIRS); do \
		$(MAKE) $(SILENT)  -C $$dir debug; \
	done

release:
	@for dir in $(SUBDIRS); do \
		$(MAKE) $(SILENT)  -C $$dir release; \
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
