package main

import (
	"log"
	"github.com/gomatic/go-vbuild"
)

//
func main() {
	log.Printf("Go toolchain build extender v%s", build.Version.String())
	build.Delegate("build")
}
