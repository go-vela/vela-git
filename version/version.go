// SPDX-License-Identifier: Apache-2.0

package version

import (
	"runtime"

	"github.com/Masterminds/semver/v3"
	"github.com/sirupsen/logrus"

	"github.com/go-vela/server/version"
)

var (
	// Arch represents the architecture information for the package.
	Arch = runtime.GOARCH
	// Commit represents the git commit information for the package.
	Commit string
	// Compiler represents the compiler information for the package.
	Compiler = runtime.Compiler
	// Date represents the build date information for the package.
	Date string
	// Go represents the golang version information for the package.
	Go string
	// OS represents the operating system information for the package.
	OS = runtime.GOOS
	// Tag represents the git tag information for the package.
	Tag = "v0.0.0"
)

// New creates a new version object for Vela that is used throughout the application.
func New() *version.Version {
	v, err := semver.NewVersion(Tag)
	if err != nil {
		logrus.Warningf("unable to parse semantic version for %s: %v", Tag, err)
	}

	return &version.Version{
		Canonical:  Tag,
		Major:      v.Major(),
		Minor:      v.Minor(),
		Patch:      v.Patch(),
		PreRelease: v.Prerelease(),
		Metadata: version.Metadata{
			Architecture:    Arch,
			BuildDate:       Date,
			Compiler:        Compiler,
			GitCommit:       Commit,
			GoVersion:       Go,
			OperatingSystem: OS,
		},
	}
}
