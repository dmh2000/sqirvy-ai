.PHONY: build test clean

SUBDIRS = cmd examples pkg web
PKG_SOURCES := $(shell find pkg -type f -name '*.go')
CMD_SOURCES := $(shell find cmd -type f -name '*.go')
SOURCES:= $(PKG_SOURCES) $(CMD_SOURCES)

# silence make output. remove -s to see make output
export SILENT=-s

debug:
	@for dir in $(SUBDIRS); do \
		$(MAKE) $(SILENT) -C $$dir debug; \
	done

release:
	@for dir in $(SUBDIRS); do \
		$(MAKE) $(SILENT) -C $$dir release; \
	done

test: debug
	@for dir in $(SUBDIRS); do \
		$(MAKE)  $(SILENT) -C $$dir test; \
	done
	@echo "Tests passed"

clean:
	@for dir in $(SUBDIRS); do \
		$(MAKE)  $(SILENT)  -C $$dir clean; \
	done
	-rm -rf bin

review:	debug
	bin/sqirvy-review -m claude-3-5-haiku-latest  $(SOURCES) >REVIEW.md

deploy: clean release test review
	git add .
	# git commit -m "Auto commit : clean, build, test, review"
