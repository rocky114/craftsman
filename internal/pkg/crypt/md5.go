package crypt

import (
	"crypto/md5"
	"fmt"
)

func GetMd5Str(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
