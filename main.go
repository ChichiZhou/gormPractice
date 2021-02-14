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
	// 如果数据库中没有这个表，则会帮你创建，如果有，则会把你创建的 model 和 数据库中的表进行更新同步
	db.AutoMigrate(&Video{})
	// 插入Record
	// 创建一个 video 实例
	//v1 := Video{1,"Pokemon", 10}
	// 将这个 video 实例插入到表中
	//db.Create(&v1)
	// 查找这个 video 表中的所有record
	//v1 := Video{Videoname: "LPL", Price:23}
	//v2 := Video{Videoname: "LDL"}   // 如果只想传某个值，则需要在前面加上这个值对应的名字
	//v3 := Video{Videoname: "Pokemon", Price:100}
	//db.Create(&v1)
	//db.Create(&v2)
	//db.Create(&v3)

	// 查询
	// 查询第一条
	video_first := new(Video)   // new 返回的是一个指针，make 返回的是一个内存地址
	db.First(&video_first)
	fmt.Println(video_first.Videoname)
	// 查询所有record
	var video []Video
	db.Find(&video)
	for _,value := range video{
		fmt.Println(value.Videoname)
	}

	// 更新
	video_update := new(Video)
	db.Where("Videoname=?", "Pokemon").Find(&video_update)
	db.Model(&video_update).Update("Price", 1000)
	// 根据传入的struct更新
	video_input := map[string]interface{}{
		"Price": 1001,
	}
	// 只更新 video_update 中的 Price 字段
	db.Model(&video_update).Select("Price").Update(video_input)

	// 删除操作
	// 如果要传递一个实例，则这个实例必须有主键，否则会删除所有实例
	db.Debug().Where("Videoname=?", "LDL").Delete(Video{})









}
