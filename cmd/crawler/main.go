package main

import (
	"context"
	"fmt"

	"github.com/rocky114/craftsman/internal/crawler"

	_ "github.com/rocky114/craftsman/internal/bootstrap"
)

func main() {
	ctx := context.Background()
	fmt.Println(crawler.Crawl(ctx, "4132010284"))
}
