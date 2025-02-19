.PHONY: debug release test clean

SUBDIRS = sqirvy util

debug:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir debug; \
	done

release:
	@for dir in $(SUBDIRS); do \
		$(MAKE) -s  -C $$dir release; \
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