-include .env
export

all: build
.PHONY: run

build: okserver.go
	go build -o build/okserver okserver.go
run:
	go run okserver.go
install:
	mkdir -p /opt/okserver
	cp build/okserver /usr/bin
	cp -r www /opt/okserver
	if [ ! -e /opt/okserver/env ]; then cp deployment/example.env /opt/okserver/env; fi
	adduser --system --no-create-home --group okserver
	chown -R okserver:okserver /opt/okserver
