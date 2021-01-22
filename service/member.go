package service

import (
	"craftsman/model"
)

func GetMembers() []map[string]interface{} {
	var items []map[string]interface{}
	model.MysqlConn.Model(&model.Member{}).Select("username", "id").Find(&items)

	return items
}

func CreateMember(member *model.Member) (err error) {
	result := model.MysqlConn.Create(member)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
