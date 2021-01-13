package model

import (
	"craftsman/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type IdModel struct {
	ID uint `gorm:"primarykey"`
}

type TimeModel struct {
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime"`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime"`
}

var MysqlConn *gorm.DB

func Bootstrap() {
	m := config.GlobalConfig.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Options

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("mysql connect failed %s", err))
	}

	MysqlConn = dbConn
}
