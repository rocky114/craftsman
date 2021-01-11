package model

import (
	"time"
)

type IdModel struct {
	ID uint `gorm:"primarykey"`
}

type TimeModel struct {
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime"`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime"`
}

