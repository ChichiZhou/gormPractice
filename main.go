package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Video struct {
	Videoid int
	Videoname string
	Price int
}
func main() {
	fmt.Println("hello")
	db, _ := gorm.Open("mysql", "root:zhouchichi@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	// 通过 gorm 创建 video 这个表
	//db.AutoMigrate(&Video{})
	// 创建一个 video 实例
	//v1 := Video{1,"Pokemon", 10}
	// 将这个 video 实例插入到表中
	//db.Create(v1)
	// 查找这个 video 表中的所有record
	var video []Video
	db.Find(&video)
	for _,value := range video{
		fmt.Println(value.Videoname)
	}
}
