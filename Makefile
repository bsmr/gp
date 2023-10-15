# === Makefile for less PIA ===

VERSION=$(shell git branch --show-current)-$(shell date --utc '+%Y%m%d-%H%M%S')

all: clean build

build:	bin/gp

bin/gp:
	go build -ldflags "-X github.com/bsmr/gp/internal/version.VersionText=$(VERSION)" -o bin/gp cmd/gp/main.go

clean:
	find bin -type f -delete

clobber:
	rm -rf bin

# === End Of File ===