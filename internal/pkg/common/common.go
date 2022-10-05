package common

import (
	"os"
	"path/filepath"
)

var rootDir string

func init() {
	rootDir, _ = os.Getwd()
	rootDir = filepath.Dir(filepath.Dir(rootDir))
}

func GetRootDir() string {
	return rootDir
}
