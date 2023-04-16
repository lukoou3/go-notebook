package version

import (
	"fmt"
)

const (
	ServiceName = "go-notebook"
)

var (
	GIT_TAG    string
	GIT_COMMIT string
	GIT_BRANCH string
	BUILD_TIME string
	GO_VERSION string
)

// FullVersion show the version info
func FullVersion() string {
	version := fmt.Sprintf("Version   : %s\nBuild Time: %s\nGit Branch: %s\nGit Commit: %s\nGo Version: %s\n", GIT_TAG, BUILD_TIME, GIT_BRANCH, GIT_COMMIT, GO_VERSION)
	return version
}

// Short 版本缩写
func Short() string {
	return fmt.Sprintf("%s[%s %s]", GIT_TAG, BUILD_TIME, GIT_COMMIT)
}

func init() {
	GIT_TAG = "1"
	GIT_COMMIT = "2"
	GIT_BRANCH = "3"
	BUILD_TIME = "4"
	GO_VERSION = "5"
}
