-include .env
export

all: build
.PHONY: run

build: main.go
	go build -o bin/oksi-http main.go
run:
	go run main.go
install:
	go build -o bin/oksi-http main.go
	cp bin/oksi-http /usr/bin