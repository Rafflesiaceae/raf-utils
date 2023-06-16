.PHONY: *
SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eo pipefail -c
.SILENT:

VERSION :=
SOME_ENV_VAR := ${SOME_ENV_VAR}

help:
	awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } /^# \{\{\{/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 7) } ' $(MAKEFILE_LIST)

# {{{ Build
build: ## Build raf-utils
	go build

test: ## Build and run some basic tests
	set -x
	make -s build
	
	# test retab
	cp "./tests/example_retab_input" "./tests/example_retab_output"
	./raf-utils retab "./tests/example_retab_output" 4 2 2
	diff "./tests/example_retab_output" "./tests/example_retab_expected_output"
