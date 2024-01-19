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

@prisma *arg:
    go run github.com/steebchen/prisma-client-go $@

gen:
	@just prisma generate

push:
	@just prisma db push

seed:
	go run ./scripts/seed

css:
	npx --yes tailwindcss -i public/_base.css -o public/styles.css
