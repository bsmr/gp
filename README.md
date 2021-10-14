# gp

Go Package - a tool to create go packages

## Abstract

- Create Go _packages_ with a source code file and a test code file.
- Additionally omit the directory hierachy for a top-level package.
- Additionally create a command (a.k.a.) main package.
- Additionally add a structure with a _New()_, and _String()_ function.

## Installation

```shell
go install github.com/bsmr/gp/cmd/gp@latest
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
