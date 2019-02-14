.PHONY: build lint golint help

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
