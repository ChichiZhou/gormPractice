package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Video struct {
	Videoid int `gorm:"primary_key;auto_increment"`
	Videoname string
	Price int  `gorm:"default:'12'"`    // 如果添加默认值，则需要把原来的表删掉再运行一遍 AutoMigrate
}
func main() {
	fmt.Println("hello")
	db, _ := gorm.Open("mysql", "root:zhouchichi@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	// 通过 gorm 创建 video 这个表
	db.AutoMigrate(&Video{})
	// 创建一个 video 实例
	//v1 := Video{1,"Pokemon", 10}
	// 将这个 video 实例插入到表中
	//db.Create(&v1)
	// 查找这个 video 表中的所有record
	//v1 := Video{Videoname: "LPL", Price:23}
	//db.Create(&v1)
	v2 := Video{Videoname: "LDL"}   // 如果只想传某个值，则需要在前面加上这个值对应的名字

	db.Create(&v2)
	var video []Video
	db.Find(&video)
	for _,value := range video{
		fmt.Println(value.Videoname)
	}
}
