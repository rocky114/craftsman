package main

import (
	"context"
	"fmt"

	_ "github.com/rocky114/craftsman/internal/bootstrap"
	"github.com/rocky114/craftsman/internal/crawler"
)

func main() {
	ctx := context.Background()
	fmt.Println(crawler.Crawl(ctx, "4132011463", "2022"))
}
