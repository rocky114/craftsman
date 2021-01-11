package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(3 * time.Second)

	fmt.Println("当前时间为:", time.Now())

	go func() {
		for {
			t := <-ticker.C
			fmt.Println("当前时间为：", t)
		}
	}()

	for {
		time.Sleep(time.Second * 1)
	}
}
