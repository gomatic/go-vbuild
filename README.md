# go-vbuild

A Go subcommand [extender](//github.com/gomatic/extender) that adds `-ldflags -X` variables to
`go build` and `go install`.

[![reportcard](https://goreportcard.com/badge/github.com/gomatic/go-vbuild)](https://goreportcard.com/report/github.com/gomatic/go-vbuild)
[![build](https://travis-ci.org/gomatic/go-vbuild.svg?branch=master)](https://travis-ci.org/gomatic/go-vbuild)

It runs:

	go (build|install) [args...] -ldflags \
	  -X github.com/gomatic/go-vbuild.Who=${USER} \
	  -X github.com/gomatic/go-vbuild.Where=${HOST} \
	  -X github.com/gomatic/go-vbuild.Patch=$(git show -s --format=%ct)-$(git log --pretty=format:'%h' -n 1)

# Example

See [cmds/go-versioning](cmds/go-versioning/main.go)

# Installation

:warning: This installs `${GOBIN}/go-build`, `${GOBIN}/go-install`. These are not intended to be
executed directly. They are utilized by [extender](//github.com/gomatic/extender) to override the
`go build` and `go install` commands.

    go get github.com/gomatic/go-vbuild

