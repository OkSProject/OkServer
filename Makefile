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
	useradd -m okserver -s /bin/bash
	mkdir -p /home/okserver/http/
	chmod -R +wx /home/okserver/http/
	cp bin/okserver /usr/bin
	cp -r ./ /home/okserver/http/
	if [ ! -e /home/okserver/http/env ]; then cp deployment/example.env /home/okserver/http/env; fi