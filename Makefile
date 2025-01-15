.PHONY: build test clean

SUBDIRS = cmd web pkg/api cmd/sqirvy-review

# silence make output. remove -s to see make output
export SILENT=-s

build:
	@for dir in $(SUBDIRS); do \
		$(MAKE) $(SILENT) -C $$dir build; \
	done

test:
	@for dir in $(SUBDIRS); do \
		$(MAKE)  $(SILENT) -C $$dir test; \
	done
	@echo "Tests passed"

clean:
	@for dir in $(SUBDIRS); do \
		$(MAKE)  $(SILENT)  -C $$dir clean; \
	done
	-rm -rf bin
