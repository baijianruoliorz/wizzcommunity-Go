package logic

import (
	"wizzcommunity/dao/mysql"
	"wizzcommunity/models"
	"wizzcommunity/pkg/jwt"
	"wizzcommunity/pkg/snowflake"
)

/*
*  @author liqiqiorz
*  @data 2020/10/20 15:13
 */
//这里写放业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	//	判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//	生成UUID
	userID := snowflake.GenID()
	//    构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//	传递的是指针,就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token

	return
}
