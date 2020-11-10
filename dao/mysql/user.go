package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"wizzcommunity/models"
)

/*
*  @author liqiqiorz
*  @data 2020/10/20 15:13
 */
//这一层用来和数据库交互
const secret = "gaoxiaoxiandyangxiangrui.cp"

//CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}
func SelectUserExist() (err error) {
	sqlStr := `select count(*) from user`
	var count int64
	db.Get(&count, sqlStr)
	fmt.Println("已连接")
	fmt.Println(count)
	return
}

//InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//	对密码进行加密
	user.Password = encryptPassword(user.Password)
	//	执行sql语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)

	return
}

//加密方法
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)

	return
}
