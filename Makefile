PROJECT_NAME := go-form3
PKG := "github.com/jonatak/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all dep lint unit-test integration-test

dep: ## Get dependencies
	@go get golang.org/x/lint/golint
	@go get -v -d ./...

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

unit-test: ## Run only unit test
	@go test -v -cover --tags=unit ${PKG_LIST}

integration-test: ## Run integration test
	@docker-compose up -d -V postgresql vault accountapi
	@docker-compose run integration-test
	@docker-compose down -v

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
