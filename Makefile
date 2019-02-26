.PHONY: test coverage build lint golint help

TMP:=.tmp
MAKE_TMP:=$(shell mkdir -p $(TMP))

COVERAGE_FILE:=$(TMP)/test-coverage.txt

#? test: run tests
test:
	go test $(shell go list ./... | grep -v /.tmp/ ) -v -coverprofile ${COVERAGE_FILE}

#? coverage: run tests with coverage report
coverage: test
	go tool cover -func=${COVERAGE_FILE}

#? build: compile binary
build:
	CGO_ENABLED=0 go build -ldflags '-w -s' -o bin/cphalo .

#? lint: run a meta linter
lint:
	@hash golangci-lint || (echo "Download golangci-lint from https://github.com/golangci/golangci-lint#install" && exit 1)
	golangci-lint run

#? golint: run golint
golint:
	@golint -set_exit_status

#? help: display help
help: Makefile
	@printf "Available make targets:\n\n"
	@sed -n 's/^#?//p' $< | column -t -s ':' |  sed -e 's/^/ /'
