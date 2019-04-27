TESTARGS=-v -race -cover

ifeq ($(TRAVIS), true)
TESTARGS=-v -race -coverprofile=coverage.txt -covermode=atomic
endif

.PHONY: dep
## Fetch dependencies
dep:
	@go get -u golang.org/x/tools/cmd/goimports
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: check
## Run checks against the codebase
check:
	@golint -set_exit_status .
	@go vet ./...
	@goimports -l . | tee /dev/tty | xargs -I {} test -z {}
	@golangci-lint run

.PHONY: doc
## Start a local godoc instance
doc:
	@echo "godoc starting - you can check the documentation at http://localhost:6060/pkg/github.com/stefanoj3/warp/"
	@godoc -http=":6060"

.PHONY: test
## Execute tests
test:
	@echo "Executing tests"
	go test $(TESTARGS) ./...

.PHONY: fix
## Run goimports against the source code
fix:
	@goimports -w .

.PHONY: help
## Display this help screen - requires gawk
help:
	@gawk 'match($$0, /^## (.*)/, a) \
		{ getline x; x = gensub(/(.+:) .+/, "\\1", "g", x) ; \
		printf "\033[36m%-30s\033[0m %s\n", x, a[1]; }' $(MAKEFILE_LIST) | sort