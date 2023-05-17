package main

import (
	"fmt"

	_ "github.com/rocky114/craftsman/internal/bootstrap"
	"github.com/rocky114/craftsman/internal/service/crawler"
)

func main() {
	if err := crawler.ScrapeAdmissionMajorScoreSuda(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
}
