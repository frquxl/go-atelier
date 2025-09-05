package cmd

import (
	"runtime/debug"
)

// Version holds the application's version string.
// It's initialized in the init() function by reading build information.
var Version = "(devel)" // Default value

func init() {
	if info, ok := debug.ReadBuildInfo(); ok {
		// The module version is set by the Go linker using VCS information.
		// For a release, this will be the git tag.
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}