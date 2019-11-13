.PHONY: default
default: build run

.PHONY: build
build: ## Builds the app binary.
	GOOS=linux GO111MODULE=on go build -o eather cmd/*.go

.PHONY: run
run: ## Runs the app on localhost, port 8000.
	./eather server

.PHONY: format
format: ## Runs the code formatter.
	gofumpt -s -w .
	gofumports -w .

.PHONY: count-lines
count-lines: ## Prints report with number of Go code lines.
	@find . -name "*.go" | xargs wc -l

.PHONY: help
help: ## Displays all targets with descs.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
