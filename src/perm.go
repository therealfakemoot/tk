package perm

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	logrus "github.com/Sirupsen/logrus"
	kingpin "github.com/alecthomas/kingpin"
}

// Hints for syntax/flags

// Instantiate logrus for logging
var logger = logrus.New()

// Permissions are defined as uint32 strings
type Permissions uint32

// Import permissions ruleset(s)
// - Unix/Linux
// - cPanel/WHM

// Show current target's permissions and pending changes
func Stat(target string) (Permissions, error) {
	file := target
	info, _ := os.Stat(file)
	mode := info.Mode()
}

// Set permissions from ruleset(s) on target

// Execution
func main() {}
