package main

import (
	"fmt"

	_ "github.com/rocky114/craftsman/internal/bootstrap"
	"github.com/rocky114/craftsman/internal/service/scraper"
)

func main() {
	if err := scraper.ScrapeAdmissionMajorScoreSuda(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
}
