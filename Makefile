.DEFAULT_GOAL := all

.PHONY: all
all: tidy build

# Build options

ROOT_PACKAGE=github.com/wxc/cmdb
VERSION_PACKAGE=github.com/marmotedu/component-base/pkg/version

# Includes

include scripts/make-rules/common.mk
#include scripts/make-rules/golang.mk

# Targets

.PHONY: build
build:
	echo $(VERSION)
	@$(MAKE) go.build

.PHONY: tidy
tidy:
	@$(GO) mod tidy
