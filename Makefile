BINARY_NAME=droxy
BINARY_FILE_PATH=".build/$(BINARY_NAME)"

setup: ## Install tools
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s v1.17.1
	mkdir .bin && mv bin/golangci-lint .bin/golangci-lint && rm -rf bin

lint: ## Run all the linters
	golangci-lint help linters
	golangci-lint run --enable=goimports --enable=gofmt --enable=gocyclo --enable=nakedret --enable=scopelint --enable=stylecheck
	
test-with-coverage: ## Run all the tests
	rm -f coverage.tmp && rm -f coverage.txt
	echo 'mode: atomic' > coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -race -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

test: ## Run all the tests
	go version
	go env
	go list ./... | xargs -n1 -I{} sh -c 'go test -race {}'

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

ci: test-with-coverage codecov build ## Run all the tests and code checks

functional-tests: build ## Runs functional tests on built binary
	cp ".build/$(BINARY_NAME)" ".test/$(BINARY_NAME)"
	cd .test && ./run.sh

local-functional-tests: build ## Runs functional tests, that does not run on drone.io
	cp ".build/$(BINARY_NAME)" ".test/$(BINARY_NAME)"
	cd .test && ./run-local.sh
		
all-functional-tests: functional-tests local-functional-tests ## Runs all functional tests
	cd .test && ./run-local.sh
	
codecov:
	codecov -t f064b312-d8a2-4f05-b5cd-f4df37dcfc89

unsafe-build: ## build binary to .build folder without testing
	rm -f $(BINARY_FILE_PATH)
	go build -o $(BINARY_FILE_PATH) main.go
	cd .droxy && ../$(BINARY_FILE_PATH) clones -f

build: ## build binary to .build folder
	rm -f $(BINARY_FILE_PATH) 
	go build -o $(BINARY_FILE_PATH) main.go

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