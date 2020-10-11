PROJECT_NAME := go-form3
PKG := "github.com/jonatak/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all lint

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}
