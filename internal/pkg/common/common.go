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

func GetLimitAndOffset(page, size int32) (int32, int32) {
	var limit int32 = 10
	if size > 0 {
		limit = size
	}
	offset := (page - 1) * limit

	return limit, offset
}
