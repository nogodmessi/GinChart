package main

import (
	"Gin+WebSocket/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:203900@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb3&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移
	db.AutoMigrate(&models.UserBasic{})
}
