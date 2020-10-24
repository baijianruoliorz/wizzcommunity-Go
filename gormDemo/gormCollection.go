package main

/*
*  @author liqiqiorz
*  @data 2020/10/25 00:07
 */
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 定义模型
type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	// 2, 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	// 3, 创建
	//u1 := User{Name: "eryajf", Age: 20}
	//db.Create(&u1)
	//u2 := User{Name: "jinzhu", Age: 22}
	//db.Create(&u2)

	// 4,查询
	var user []User
	db.Debug().First(&user) // SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("根据主键查询第一条记录：", user)

	db.Debug().Take(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL LIMIT 1
	fmt.Println("随机获取一条记录：", user)

	db.Debug().Last(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` DESC LIMIT 1
	fmt.Println("根据主键查询最后一条记录：", user)

	db.Debug().Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL
	fmt.Println("查询所有的记录：", user)

	db.Debug().First(&user, 2)      //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`id` = 2)) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询指定的某条记录：", user) //仅当主键为整型时可用

	db.Debug().Where("name = ?", "jinzhu").First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询第一条匹配条件记录：", user)

	db.Debug().Where("name = ?", "jinzhu").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu'))
	fmt.Println("查询所有匹配条件的记录：", user)

	db.Debug().Where("name <> ?", "jinzhu").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name <> 'jinzhu'))
	fmt.Println("查询name不等于jinzhu的所有记录：", user)

	db.Debug().Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name IN ('jinzhu','jinzhu 2')))
	fmt.Println("查询name在jinzhu和jinzhu 2的所有记录：", user)

	db.Debug().Where("name LIKE ?", "%jin%").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name LIKE '%jin%'))
	fmt.Println("查询name包含jin的所有记录：", user)

	db.Debug().Where("name = ? AND age >= ?", "jinzhu", "20").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu' AND age >= '20'))
	fmt.Println("查询两个条件都符合的所有记录：", user)

	oneDay, _ := time.ParseDuration("-24h")
	lastWeek := time.Now().Add(oneDay * 7)
	db.Debug().Where("updated_at > ?", lastWeek).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((updated_at > '2020-03-01 19:45:11'))
	fmt.Println("查询一周内更新的用户记录：", user)

	today := time.Now()
	db.Debug().Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((created_at BETWEEN '2020-03-01 19:52:51' AND '2020-03-08 19:52:51'))
	fmt.Println("查询一周内创建的记录：", user)

	db.Debug().Where(&User{Name: "jinzhu", Age: 22}).First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu') AND (`user`.`age` = 22)) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("通过结构体查询：", user)

	db.Debug().Where(map[string]interface{}{"name": "jinzhu", "age": 22}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu') AND (`user`.`age` = 22))
	fmt.Println("通过map查询：", user)

	db.Debug().Where([]int64{1, 2}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`id` IN (1,2)))
	fmt.Println("通过主键的切片查询：", user)

	db.Debug().Not("name", "jinzhu").First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` NOT IN ('jinzhu'))) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询name不是jinzhu的第一条记录：", user)

	db.Debug().Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` NOT IN ('jinzhu','jinzhu 2')))
	fmt.Println("查询name不在jinzhu或jinzhu2的所有记录：", user)

	db.Debug().Not([]int64{1, 2, 3}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`id` NOT IN (1,2,3)))
	fmt.Println("查询主键不是1，2，3的所有记录：", user)

	db.Debug().Not([]int64{}).First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询所有用户中的第一个：", user)

	db.Debug().Not("name = ?", "jinzhu").First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND (NOT (name = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询name不是jinzhu的第一个用户：", user)

	db.Debug().Not(User{Name: "jinzhu"}).First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` <> 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("通过结构体查询name不是jinzhu的第一个用户：", user)

	db.Debug().Where("age > ?", 25).Or("age < ?", 23).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((age > 25) OR (age < 23))
	fmt.Println("查询年龄小于23的或者大于25的所有记录：", user)

	// struct
	db.Debug().Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu') OR (`user`.`name` = 'jinzhu 2'))
	fmt.Println("结构体：查询名字是jinzhu的或者是jinzhu 2的所有记录：", user)

	// map
	db.Debug().Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu') OR (`user`.`name` = 'jinzhu 2'))
	fmt.Println("map：查询名字是jinzhu的或者是jinzhu 2的所有记录：", user)

	db.Debug().First(&user, 3)          //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`id` = 3)) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("根据主键查询指定的某条记录：", user) //仅当主键为整型时可用

	db.Debug().First(&user, "id = ?", "string_primary_key") //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((id = 'string_primary_key')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("根据主键是非整形主键获取记录：", user)

	db.Debug().Find(&user, "name = ?", "jinzhu") //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu'))
	fmt.Println("查询name为jinzhu的记录：", user)

	db.Debug().Find(&user, "name <> ? AND age > ? ", "jinzhu", "20") //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name <> 'jinzhu' AND age > '20' ))
	fmt.Println("查询name不是jinzhu且年龄大于20的记录：", user)

	db.Debug().Find(&user, User{Age: 20}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`age` = 20))
	fmt.Println("通过结构体查询年龄是20的所有记录：", user)

	db.Debug().Find(&user, map[string]interface{}{"age": 20}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`age` = 20))
	fmt.Println("通过map查询年龄是20的所有记录：", user)

	db.Debug().FirstOrInit(&user, User{Name: "non_existing"}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询name为non_existing的记录：", user)

	db.Debug().Where(User{Name: "jinzhu"}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("通过结构体查询name为jinzhu的记录：", user)

	db.Debug().FirstOrInit(&user, map[string]interface{}{"name": "jinzhu"}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("通过map查询name为jinzhu的记录：", user)

	// 未找到
	db.Debug().Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "non_existing"}).Attrs("age", 20).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	// 找到
	db.Debug().Where(User{Name: "jinzhu"}).Attrs(User{Age: 50}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	// 找到
	db.Debug().Where(User{Name: "jinzhu"}).Assign(User{Age: 50}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().FirstOrCreate(&user, User{Name: "non_existing"}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "jinzhu"}).FirstOrCreate(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "jinzhu"}).Attrs(User{Age: 30}).FirstOrCreate(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Select("name", "age").Find(&user) //SELECT name FROM `user`  WHERE `user`.`deleted_at` IS NULL'age'
	fmt.Println("查询表中name字段参数为age的记录：", user)

	db.Debug().Select([]string{"name", "age"}).Find(&user) //SELECT name, age FROM `user`  WHERE `user`.`deleted_at` IS NULL
	fmt.Println("列出表中name与age字段：", user)

	db.Debug().Order("age desc,name").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY age desc,name
	fmt.Println("根据年龄排序来查询：", user)

	db.Debug().Order("age desc").Order("name").Find(&user)
	fmt.Println("根据多个条件排序查询：", user)

}
