##
# SEC company valuation
#
# @file
# @version 0.1

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-cli:
	go build -o bin/sec-valuation cmd/cli/cli.go

build-server:
	docker build -t ahewins/sec-company-valuation -f ./server.Dockerfile

.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u github.com/alecthomas/gometalinter
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/dep/cmd/dep
	gometalinter --install --update
	@$(MAKE) dep

.PHONY: clean
clean:
	go clean

.DEFAULT_GOAL := build
# end
