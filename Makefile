# ------------------------------------------------------
# VARIABLES
# ------------------------------------------------------
SHELL			:= /bin/bash -o pipefail

# ------------------------------------------------------
# HELP
# ------------------------------------------------------
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# ------------------------------------------------------
# MAIN COMMANDS
# ------------------------------------------------------
.PHONY: all
all: gen format test

.PHONY: gen
gen: ## Generate code
	rm -f $$(find . -type f -name "*mock_test.go")
	go generate ./...
	go mod tidy -go=1.19 -compat=1.19

.PHONY: format
format: ## Format the source code
	go run golang.org/x/tools/cmd/goimports -w .

.PHONY: test
test: ## Run tests
	go test -count=1 -timeout 10m -v -cover -race ./...
