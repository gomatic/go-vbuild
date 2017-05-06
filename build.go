package build

import (
	"fmt"
	"sync"
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
