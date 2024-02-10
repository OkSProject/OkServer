-include .env
export

all: build
.PHONY: run

build: main.go
	go build -o bin/okserver main.go
run:
	go run main.go
install:
	go build -o bin/okserver main.go
	cp bin/okserver /usr/bin