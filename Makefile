export GO111MODULE=on
BINARY_NAME=droxy
BINARY_FILE_PATH=".build/$(BINARY_NAME)"

setup: ## Install all the build and lint dependencies
	wget -O- https://git.io/vp6lP | sh 
	go get -u golang.org/x/tools/cmd/goimports

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

lint: ## Run all the linters
	gometalinter --vendor --disable-all \
		--enable=deadcode \
		--enable=gocyclo \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=staticcheck \
		--enable=gofmt \
		--enable=golint \
		--enable=goimports \
		--enable=dupl \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--enable=varcheck \
		--enable=structcheck \
		--enable=interfacer \
		--enable=goconst \
		--deadline=10m \
		./... | grep -v "mocks"

ci: test-with-coverage codecov build ## Run all the tests and code checks

functional-tests: build ## Runs functional tests on built binary
	cp ".build/$(BINARY_NAME)" ".test/$(BINARY_NAME)"
	cd .test && ./run.sh

local-functional-tests:  ## Runs functional tests, that does not run on drone.io
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
	CGO_ENABLED=0 go build -o $(BINARY_FILE_PATH) main.go

install: build ## build with tests, then install to <gopath>/src
	cp $(BINARY_FILE_PATH) $$GOPATH/bin/$(BINARY_NAME)

build-release: ## builds the checked out version into the .release/${tag} folder
	.release/build.sh

build-release-test: ## builds the checked out version into the .release/${tag} folder
	.release/build.sh test

# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help