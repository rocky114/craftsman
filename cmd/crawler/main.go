package main

import (
	"fmt"

	_ "github.com/rocky114/craftsman/internal/bootstrap"
	"github.com/rocky114/craftsman/internal/pkg/path"
)

func main() {
	//pathGetCurrentPath()

	fmt.Println(path.GetRootPathByCaller())
	fmt.Println(path.GetRootPath())
}
