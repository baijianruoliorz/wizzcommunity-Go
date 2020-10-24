package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
*  @author liqiqiorz
*  @data 2020/10/24 23:42
 */
//会自动创建哒,命名规则是user_infos
//不过数据库要自己来哦!
type UserInfo struct {
	//这里都会变成小写!
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//	自动迁移
	db.AutoMigrate(&UserInfo{})
	//u1:=UserInfo{1,"liqiqiorz","男","写代码"}
	//u2:=UserInfo{2,"baijianruoliorz","女","谈恋爱"}
	//创建记录
	//db.Create(&u1)
	//db.Create(&u2)
	//查询
	var u = new(UserInfo)
	//找出第一个  &main.UserInfo{ID:0x1, Name:"liqiqiorz", Gender:"男", Hobby:"写代码"}
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	//查询
	db.Find(&uu, "hobby=?", "写代码")
	fmt.Printf("%#v\n", uu)
	//更新
	db.Model(&u).Update("hobby", "搞基")
	//删除
	db.Delete(&u)

}
