# Arthur Mingard

#
# Variables
#
USER=mingard
PACKAGE=sitemap-format
VERSION=0.0.2

# Go command variables
GOCMD=go
GOTEST=$(GOCMD) test


# Compiler variables
SRC ?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKGS = $(shell go list ./... | grep -v /vendor)

.PHONY: version test lint fmt update-pkg-cache

%:
	@true

version:
	@echo $(VERSION)

test:
	$(GOTEST) $(PKGS)

lint: $(GOMETALINTER)
	$(GOMETALINTER) ./... --vendor --fast --disable=maligned

$(GOMETALINTER):
	$(GOGET) -u github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install 1>/dev/null

fmt:
	gofmt -l -w $(SRC)

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/$(USER)/$(PACKAGE)
#	go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)