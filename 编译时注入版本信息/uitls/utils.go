package utils

import (
	"fmt"
	"runtime"
)

var (
	version      string
	gitBranch    string
	gitTag       string
	gitCommit    string
	gitTreeState string
	buildDate    string
)

type Info struct {
	Version      string `json:"version"`
	GitBranch    string `json:"gitBranch"`
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	return info.GitCommit
}

func GetVersion() Info {
	return Info{
		Version:      version,
		GitBranch:    gitBranch,
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
