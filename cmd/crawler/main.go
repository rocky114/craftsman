package main

import (
	_ "github.com/rocky114/craftsman/internal/bootstrap"
	"github.com/rocky114/craftsman/internal/crawler"
)

func main() {
	crawler.Selenium()
	//ctx := context.Background()
	//fmt.Println(crawler.Crawl(ctx, "4132011117", "2022"))
}
