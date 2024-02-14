-include .env
export

all: build
.PHONY: run

build: oksi-http.go
	go build -o bin/oksi-http oksi-http.go
run:
	go run oksi-http.go
install:
	go build -o bin/oksi-http oksi-http.go
	cp bin/oksi-http /usr/bin