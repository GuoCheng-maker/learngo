package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	s := uuid.NewString()
	fmt.Println(s)
	db, _ := gorm.Open("mysql", "root:@(127.0.0.1:3306)/studygo?charset=utf8mb4&parseTime=True&loc=Local")
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// 创建数据表
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{ID: 2, Name: "马云", Gender: "男", Hobby: "挣钱"}
	log.Println("开始创建用户信息")
	result := db.Create(&u1)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	log.Println("用户信息创建完毕")

}
