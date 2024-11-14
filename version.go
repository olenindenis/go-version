package version

import (
	"fmt"
	"runtime"
	"time"
)

var (
	gitVersion = "v0.0.0"
	gitCommit  = ""
	buildDate  = time.RFC3339Nano
)

type info struct {
	GitVersion string `json:"git_version" yaml:"git_version"`
	GitCommit  string `json:"git_commit" yaml:"git_commit"`
	BuildDate  string `json:"build_date" yaml:"build_date"`
	GoVersion  string `json:"go_version" yaml:"go_version"`
	Compiler   string `json:"compiler" yaml:"compiler"`
	Platform   string `json:"platform" yaml:"platform"`
}

func Info() info {
	return info{
		GitVersion: gitVersion,
		GitCommit:  gitCommit,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
