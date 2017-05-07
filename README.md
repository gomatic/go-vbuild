# go-vbuild

A Go subcommand [extender](//github.com/gomatic/extender) that adds `-ldflags -X` variables to
`go build` and `go install`.

[![reportcard](https://goreportcard.com/badge/github.com/gomatic/go-vbuild)](https://goreportcard.com/report/github.com/gomatic/go-vbuild)
[![build](https://travis-ci.org/gomatic/go-vbuild.svg?branch=master)](https://travis-ci.org/gomatic/go-vbuild)
[![godoc](https://godoc.org/github.com/gomatic/go-vbuild?status.svg)](https://godoc.org/github.com/gomatic/go-vbuild)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

It runs:

	go (build|install) [args...] -ldflags \
	  -X github.com/gomatic/go-vbuild.Who=${USER} \
	  -X github.com/gomatic/go-vbuild.Where=${HOST} \
	  -X github.com/gomatic/go-vbuild.Patch=$(git show -s --format=%ct)-$(git log --pretty=format:'%h' -n 1)

# Example

See [cmds/go-versioning](cmds/go-versioning/main.go)

# Installation

:warning: This installs `${GOBIN}/go-build`, `${GOBIN}/go-install`. These are intended to be
executed by [extender](//github.com/gomatic/extender) to override the
`go build` and `go install` commands.

    go get github.com/gomatic/go-vbuild

