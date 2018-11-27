.PHONY: help
.DEFAULT_GOAL=help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## run the example
	go run main.go

info: ## show dependencies
	go mod graph

download: ## download modules to cache
	go mod download

install: ## build a binary in $GOHOME/bin
	go install main.go