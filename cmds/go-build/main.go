package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"

	"strings"

	"github.com/gomatic/commander"
	"github.com/gomatic/go-vbuild"
)

var ws = regexp.MustCompile(`[[:space:][:punct:]]`)

//
func main() {
	log.Println(build.Version.Detailed())

	username := ""
	usr, err := user.Current()
	if err == nil {
		username = usr.Username
	}

	hostname, err := os.Hostname()

	name, exists := os.LookupEnv("VERSIONING_STRUCT")
	if !exists {
		name = "github.com/gomatic/go-build"
	}
	X := fmt.Sprintf(`-X %s.%%s=%%s`, name)

	args := []string{}
	args = append(args, fmt.Sprintf(X, "who", ws.ReplaceAllString(username, "_")))
	args = append(args, fmt.Sprintf(X, "where", ws.ReplaceAllString(hostname, "_")))

	// TODO climb parents looking for repository
	if _, err := os.Stat(".hg"); err == nil {
		hg()
	} else if _, err := os.Stat(".bzr"); err == nil {
		bzr()
	} else if _, err := os.Stat(".svn"); err == nil {
		svn()
	} else { //TODO don't assume git. if _, err := os.Stat(".git"); err == nil {
		args = append(args, git(X))
	}

	goroot, exists := os.LookupEnv("GOROOT")
	if !exists {
		log.Println("Missing GOROOT")
		os.Exit(1)
	}
	cmd := commander.New("").Inherit(2)
	cmd.Binary = filepath.Join(goroot, "bin", "go")
	cmd.Args("-ldflags", strings.Join(args, " "))

	//

	err = cmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

//
func git(x string) string {
	ctm, err := exec.Command("git", "show", "-s", "--format=%ct").Output()
	if err != nil {
		log.Fatal(err)
	}
	cid, err := exec.Command("git", "log", "--pretty=format:'%h'", "-n", "1").Output()
	if err != nil {
		log.Fatal(err)
	}

	version := ws.ReplaceAllString(string(ctm), "") + "-" + ws.ReplaceAllString(string(cid), "")
	return fmt.Sprintf(x, "patch", version)

	// TODO get tag version if it exists
	//tag := majorVersion
	//fmt.Printf("-ldflags -X main.tag=%s\n", tag)
}

// TODO
func hg() string {
	return ""
}

// TODO
func svn() string {
	return ""
}

// TODO
func bzr() string {
	return ""
}
