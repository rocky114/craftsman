package main

import (
	"context"
	"fmt"
	"time"

	_ "github.com/rocky114/craftsman/internal/bootstrap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		go job(ctx, i)
	}

	time.Sleep(5 * time.Second)
}

func job(ctx context.Context, i int) {
	httpChannel := make(chan struct{})
	select {
	case <-ctx.Done():
		fmt.Printf("ctx timeout, number:%d\n", i)
	case <-httpChannel:
		fmt.Printf("job done, number:%d\n", i)
	}
}
