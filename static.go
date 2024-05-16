package dblogger

import (
	_ "embed"
	"runtime/debug"
)

//go:embed data/combined/schema.sql
var SQLSchema string

//go:embed data/combined/queries.sql
var SQLQueries string

// gitRevision is the git revision of the application
var gitRevision string

// buildInfo is the build information of the application
var buildInfo debug.BuildInfo

// gitRevision is the git revision of the application
var buildSum string

// goVersion is the version of the Go compiler used to build the application
var goVersion string

// init is called when the package is loaded
func init() {
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range buildInfo.Settings {
			if v.Key == "vcs.revision" {
				gitRevision = v.Value
				break
			}
		}
	}
	goVersion = buildInfo.GoVersion
	buildSum = buildInfo.Main.Sum
}
