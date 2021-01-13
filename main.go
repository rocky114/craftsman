package main

import (
	"craftsman/bootstrap"
	"fmt"
)

func main() {
	fmt.Println("application starting...")

	db, _ := bootstrap.MysqlConn.DB()
	defer func() {
		_ = db.Close()
	}()
}
