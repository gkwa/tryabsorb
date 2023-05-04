SOURCES := $(shell find . -name '*.go')
TARGET := ./dist/tryabsorb_darwin_amd64_v1/tryabsorb

run: tryabsorb
	./tryabsorb

tryabsorb: $(TARGET)
	cp $< $@

$(TARGET): $(SOURCES)
	gofumpt -w $<
	goreleaser build --single-target --snapshot --clean

.PHONY: clean
clean:
	rm -f tryabsorb
	rm -f $(TARGET)
	rm -rf dist
