# gp

Go Package - a tool to create go packages

## Abstract

- Create Go _packages_ with a source code file and a test code file.
- Additionally omit the directory hierachy for a top-level package.
- Additionally create a command (a.k.a.) main package.
- Additionally add a structure with a _New()_, and _String()_ function.

## Installation

```shell
# latest version
go install github.com/bsmr/gp/cmd/gp@latest

# specify version
go install github.com/bsmr/gp/cmd/gp@v0.0.2

# install from source
go install ./cmd/gp
```

## Usage

```shell
# create a new package
gp -name MODULE

# to overwrite existing files
gp -name MODULE -force

# create a top-level package
gp -name MODULE -top

# create a main program
gp -name MODULE -cmd
```
