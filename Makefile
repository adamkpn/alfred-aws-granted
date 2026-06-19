.PHONY: build test lint clean workflow icons sign

BINARY := alfred-aws-granted
CMD := ./cmd/alfred-aws-granted
VERSION ?= 1.0.1
LINKS := aws-granted-profiles aws-granted-services aws-granted-regions aws-granted-open
CODESIGN_IDENTITY ?= -
CODESIGN_FLAGS := --force --sign "$(CODESIGN_IDENTITY)"
ifeq ($(CODESIGN_IDENTITY),-)
CODESIGN_FLAGS += --timestamp=none
else
CODESIGN_FLAGS += --options runtime --timestamp
endif

icons:
	chmod +x scripts/generate-icons.sh
	./scripts/generate-icons.sh

build: icons
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BINARY)-arm64 $(CMD)
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BINARY)-amd64 $(CMD)
	lipo -create -output $(BINARY) $(BINARY)-arm64 $(BINARY)-amd64
	rm -f $(BINARY)-arm64 $(BINARY)-amd64
	chmod +x $(BINARY)
	$(MAKE) sign
	@for link in $(LINKS); do ln -sf $(BINARY) $$link; done
	chmod +x scripts/*.sh

sign:
	codesign $(CODESIGN_FLAGS) $(BINARY)

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -f $(BINARY) $(LINKS) build/*.alfredworkflow

workflow: build
	rm -rf build
	mkdir -p build
	sed 's/1\.0\.1/$(VERSION)/g' info.plist > build/info.plist
	cd build && \
		cp ../$(BINARY) . && \
		codesign $(CODESIGN_FLAGS) $(BINARY) && \
		cp ../icon.png ../region.png . && \
		cp -r ../icons . && \
		for link in $(LINKS); do ln -sf $(BINARY) $$link; done && \
		cp -r ../services . && \
		cp -r ../scripts . && \
		chmod +x scripts/*.sh && \
		zip -r "alfred-granted-v$(VERSION).alfredworkflow" info.plist icon.png region.png icons $(BINARY) $(LINKS) services scripts
	@echo "Created build/alfred-granted-v$(VERSION).alfredworkflow"
