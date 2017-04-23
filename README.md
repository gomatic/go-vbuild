# go-pre-build

A Go subcommand extender that updates the environment prior to running `go build`.

It returns:

	-ldflags -X main.who=${USER}
	-ldflags -X main.where=${HOST}
	-ldflags -X main.version=$(git show -s --format=%ct)-$(git log --pretty=format:'%h' -n 1)

Which will be sent to `go builg` by [extender](github.com/gomatic/extender) as:

    go build -ldflags '-X main.version=[time]-[commit] -X main.who=[user] -X main.where=[host]` 
