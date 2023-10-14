# === Makefile for less PIA ===

VERSION=$(shell git branch --show-current)

all: clean build

build:	bin/gp

bin/gp:
	go build -ldflags "-X github.com/bsmr/gp.VersionText=$(VERSION)" -o bin/gp cmd/gp/main.go

clean:
	find bin -type f -delete

clobber:
	rm -rf bin

# === End Of File ===