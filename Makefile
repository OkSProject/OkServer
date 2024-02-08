-include .env
export

all: build
.PHONY: run

build: main.go
	go build -o bin/okserver main.go
run:
	go run main.go
install:
	mkdir -p /opt/okserver
	cp bin/okserver /usr/bin
	cp -r www /opt/okserver
	if [ ! -e /opt/okserver/env ]; then cp deployment/example.env /opt/okserver/env; fi
	adduser --system --no-create-home --group okserver
	chown -R okserver:okserver /opt/okserver
