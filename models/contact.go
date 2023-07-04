package models

import (
	"Gin+WebSocket/utils"
	"fmt"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁
	Type     int  //对应的类型 1好友 2群组 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

// 查询好友列表
func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts) //查询userId对应的好友有哪些
	for _, v := range contacts {
		fmt.Println(v)
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	fmt.Println(users)
	return users
}
