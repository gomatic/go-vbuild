package build

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"sync"

	"github.com/gomatic/extender/extension"
	"gopkg.in/yaml.v2"
)

var major, minor uint
var who, where, patch string

//
var Version = versioning{Major: major, Minor: minor, Patch: patch, Who: who, Where: where}

//
func init() {
	if major == uint(0) {
		major = 1
		Version.Major = major
	}
	if patch == "" {
		patch = "0"
		Version.Patch = patch
	}

	debugging, exists := os.LookupEnv("DEBUGGING")
	if exists && strings.ToLower(debugging) == "true" {
		if config, err := yaml.Marshal(yaml.MapSlice{
			{"major", major},
			{"minor", minor},
			{"patch", patch},
			{"who", who},
			{"where", where},
		}); err == nil {
			fmt.Fprintln(os.Stderr, string(config))
		}
	}
}

//
type versioning struct {
	Major, Minor      uint
	Who, Where, Patch string
	x                 sync.Mutex
}

//
func (v *versioning) Update(major, minor uint) *versioning {
	v.x.Lock()
	defer v.x.Unlock()
	v.Major, v.Minor = major, minor
	return v
}

//
func (v *versioning) String() string {
	v.x.Lock()
	defer v.x.Unlock()
	return fmt.Sprintf("%d.%d.%s", v.Major, v.Minor, v.Patch)
}

//
func (v *versioning) Detailed() string {
	v.x.Lock()
	defer v.x.Unlock()
	return fmt.Sprintf("%d.%d.%s.%s/%s", v.Major, v.Minor, v.Patch, v.Who, v.Where)
}

var ws = regexp.MustCompile(`[[:space:][:punct:]]`)

//
func Delegate(subcommand string) {
	err := extension.Delegate(subcommand, "-ldflags", MustFlags())
	if err != nil {
		log.Fatal(err)
	}
}

//
func MustFlags() string {
	flags, err := Flags()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Join(flags, " ")
}

//
func Flags() (args []string, err error) {
	username := ""
	usr, err := user.Current()
	if err == nil {
		username = usr.Username
	}

	hostname, err := os.Hostname()

	name, exists := os.LookupEnv("VERSIONING_STRUCT")
	if !exists {
		name = "github.com/gomatic/go-vbuild"
	}
	X := fmt.Sprintf(`-X %s.%%s=%%s`, name)

	args = append(args, fmt.Sprintf(X, "who", ws.ReplaceAllString(username, "_")))
	args = append(args, fmt.Sprintf(X, "where", ws.ReplaceAllString(hostname, "_")))

	// TODO climb parents looking for repository

	if _, err = os.Stat(".hg"); err == nil {
		hg()
	} else if _, err = os.Stat(".bzr"); err == nil {
		bzr()
	} else if _, err = os.Stat(".svn"); err == nil {
		svn()
	} else { //TODO don't assume git. if _, err := os.Stat(".git"); err == nil {
		args = append(args, git(X))
		err = nil
	}

	return
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
