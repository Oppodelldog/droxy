BINARY_NAME=droxy
BINARY_FILE_PATH=".build/$(BINARY_NAME)"
MAIN_FILE="main.go"

setup: ## Install tools
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.3
	go get github.com/vektra/mockery/.../

lint: ## Run the linters
	golangci-lint run

test: ## Run all the tests
	go version
	go env
	go list ./... | xargs -n1 -I{} sh -c 'go test -race {}'

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

mocks:
	mockery -name CommandBuilder -dir=cmd/proxyexecution -output=cmd/mocks
	mockery -name CommandResultHandler -dir=cmd/proxyexecution -output=cmd/mocks
	mockery -name CommandRunner -dir=cmd/proxyexecution -output=cmd/mocks
	mockery -name ConfigLoader -dir=cmd/proxyexecution -output=cmd/mocks
	mockery -name ExecutableNameParser -dir=cmd/proxyexecution -output=cmd/mocks

ci: test build lint ## Run all the tests and code checks

functional-tests: build ## Runs functional tests on built binary
	cp ".build/$(BINARY_NAME)" ".test/$(BINARY_NAME)"
	cd .test && ./run.sh

local-functional-tests: build ## Runs functional tests, that does not run on drone.io
	cp ".build/$(BINARY_NAME)" ".test/$(BINARY_NAME)"
	cd .test && ./run-local.sh
		
all-functional-tests: functional-tests local-functional-tests ## Runs all functional tests
	cd .test && ./run-local.sh

unsafe-build: ## build binary to .build folder without testing
	rm -f $(BINARY_FILE_PATH)
	go build -o $(BINARY_FILE_PATH) $(MAIN_FILE)
	cd .droxy && ../$(BINARY_FILE_PATH) clones -f

build: ## build binary to .build folder
	rm -f $(BINARY_FILE_PATH) 
	go build -o $(BINARY_FILE_PATH) $(MAIN_FILE)

install: build ## build with tests, then install to <gopath>/src
	rm -f $$GOPATH/bin/$(BINARY_NAME)
	cp $(BINARY_FILE_PATH) $$GOPATH/bin/$(BINARY_NAME)

build-release: ## builds the checked out version into the .release/${tag} folder
	.release/build.sh

build-release-test: ## builds the checked out version into the .release/${tag} folder
	.release/build.sh test

# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help