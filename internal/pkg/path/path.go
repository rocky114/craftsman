package path

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

func GetRootPathByCaller() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}

	return path.Dir(path.Dir(path.Dir(path.Dir(filename))))
}

func GetRootPath() string {
	rootPath, _ := os.Getwd()
	return rootPath
}

func GetConfigFile(filename string) string {
	return fmt.Sprintf("%s/config/%s", GetRootPath(), filename)
}

func GetTmpPath(dirs ...string) string {
	subDir := ""
	if len(dirs) > 0 {
		subDir = strings.Join(dirs, "/")
	}
	return fmt.Sprintf("%s/tmp/%s", GetRootPath(), subDir)
}

func GetLogPath() string {
	return fmt.Sprintf("%s/log/", GetRootPath())
}
