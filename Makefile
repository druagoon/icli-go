SHELL := bash

GOPATH ?= $(HOME)/.go
GOBIN ?= $(GOPATH)/bin

.PHONY: init
init:
	go install github.com/spf13/cobra-cli@latest

.PHONY: build
build:
	go build -o ./bin/icli .

.PHONY: install
install:
	go build -o "$(GOBIN)/icli" .

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: cmd-%
cmd-%:
	$(SHELL) ./tools/cmd "$*"
