package common

import "os"

var rootDir string

func init() {
	rootDir, _ = os.Getwd()
}

func GetRootDir() string {
	return rootDir
}
