package main

import (
	"log"

	"github.com/gomatic/go-build"
)

//
func main() {
	log.Println(build.Version.Detailed())
}
