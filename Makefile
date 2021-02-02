#!bin/sh
define USAGE
Commands:
    run			Run main.go
	build		Build binary
    start       Run built app
endef

run:
	go mod vendor
	go run main.go

build:
	go mod vendor
	go build -o dist/account-service

start:
	./dist/account-service