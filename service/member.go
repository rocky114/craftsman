package service

import (
	"craftsman/model"
)

func GetMembers() []map[string]interface{} {
	var items []map[string]interface{}
	model.MysqlConn.Model(&model.Member{}).Select("name", "id").Find(&items)

	return items
}
