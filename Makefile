GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

GOLANGCI_VERSION = 1.41.1

lint: ## Run Linter
	@mkdir -p bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_VERSION}

	./bin/golangci-lint run

test: ## Run Tests
	go test -v -cover -race ./...

update-pkg-cache: ## Update go pkg cache with the latest version
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
		go get -v github.com/odinn1984/go-sensibo

clean: ## Clean Go Project
	go clean

help: ## Show This Help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "  ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
	}' ${MAKEFILE_LIST}
	@echo ''
