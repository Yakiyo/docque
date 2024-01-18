set positional-arguments

alias b := build
alias u := update
alias r := run

default:
  @just --list

@run *arg:
    go run . $@

test:
	go test ./...

build:
	go build -o app .

fmt:
	go fmt ./...

update:
	go get
	go mod tidy