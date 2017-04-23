package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
)

//
const majorVersion = "1.0"

//
var version = "0"

var ws = regexp.MustCompile(`[[:space:][:punct:]]`)

//
func main() {
	log.Printf("%s.%s", majorVersion, version)

	username := ""
	usr, err := user.Current()
	if err == nil {
		username = usr.Username
	}

	hostname, err := os.Hostname()

	fmt.Printf(`-ldflags -X main.who=%s%s`, ws.ReplaceAllString(username, "_"), "\n")
	fmt.Printf(`-ldflags -X main.where=%s%s`, ws.ReplaceAllString(hostname, "_"), "\n")

	if _, err := os.Stat(".git"); err == nil {
		git()
	} else if _, err := os.Stat(".hg"); err == nil {
		hg()
	} else if _, err := os.Stat(".bzr"); err == nil {
		bzr()
	} else if _, err := os.Stat(".svn"); err == nil {
		svn()
	}
}

//
func git() {
	ctm, err := exec.Command("git", "show", "-s", "--format=%ct").Output()
	if err != nil {
		log.Fatal(err)
	}
	cid, err := exec.Command("git", "log", "--pretty=format:'%h'", "-n", "1").Output()
	if err != nil {
		log.Fatal(err)
	}

	version := ws.ReplaceAllString(string(ctm), "") + "-" + ws.ReplaceAllString(string(cid), "")
	fmt.Printf("-ldflags -X main.version=%s\n", version)

	// TODO get tag version if it exists
	//tag := majorVersion
	//fmt.Printf("-ldflags -X main.tag=%s\n", tag)
}

// TODO
func hg() {

}

// TODO
func svn() {

}

// TODO
func bzr() {

}
