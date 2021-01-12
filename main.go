package main

import (
	"craftsman/bootstrap"
	"fmt"
)

func main() {
	bootstrap.Viper()

	fmt.Println(bootstrap.GlobalConfig.Mysql.Host)

	fmt.Println("starting application...")
}
