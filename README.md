# go-vbuilder

A Go subcommand extender that adds `-ldflags -X` variables to
`go build` and `go install`.

It runs:

	go (build|install) [args...] -ldflags \
	  -X github.com/gomatic/go-vbuilder.Who=${USER} \
	  -X github.com/gomatic/go-vbuilder.Where=${HOST} \
	  -X github.com/gomatic/go-vbuilder.Patch=$(git show -s --format=%ct)-$(git log --pretty=format:'%h' -n 1)

# Example

See [cmds/go-versioning](cmds/go-versioning/main.go)
