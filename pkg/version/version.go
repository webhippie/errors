package version

import (
	"fmt"

	"github.com/coreos/go-semver/semver"
)

var (
	// Compile is the compiled version
	Compile = "0.0.0"

	// Commit indicates the current commit
	Commit = "0000000"

	// Date indicates the build date
	Date = "00000000"

	// Version represents the parsed version
	Version *semver.Version
)

func init() {
	Version = semver.New(Compile)
	Version.Metadata = fmt.Sprintf("git%s.%s", Date, Commit)
}
