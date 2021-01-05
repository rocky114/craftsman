package main

import (
	"craftsman/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:rocky114@tcp(81.68.171.7:3306)/rocky?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	member := model.Member{
		Name:    "admin",
		IsAdmin: 1,
		IsSuper: 1,
		Status:  "active",
	}

	result := db.Create(&member)

	fmt.Println(result)
}
