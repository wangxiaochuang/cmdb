.DEFAULT_GOAL := all

.PHONY: all
all: tidy build

# Build options

ROOT_PACKAGE=github.com/wxc/cmdb
VERSION_PACKAGE=github.com/marmotedu/component-base/pkg/version

# Includes

include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/tools.mk

# Targets

.PHONY: build
build:
	@$(MAKE) go.build

.PHONY: lint
lint:
	@$(MAKE) go.lint

.PHONY: tidy
tidy:
	@$(GO) mod tidy

localIp := 127.0.0.1
.PHONY: run
run:
	@_output/platforms/darwin/amd64/apiserver --addrport=${localIp}:8080 --logtostderr=false --log-dir=./logs --v=3 --regdiscv=127.0.0.1:2181  --enable-auth=false
