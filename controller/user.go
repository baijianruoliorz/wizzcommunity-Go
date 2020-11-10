package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"wizzcommunity/dao/mysql"
	"wizzcommunity/logic"
	"wizzcommunity/models"
)

/*
*  @author liqiqiorz
*  @data 2020/10/20 15:13
 */

/*
//@Summary 简单注册接口
//@Description 用户注册用
//@Tags User
//@Accept application/json
//@Produce application/json
//@param user body models.ParamSignUp true "注册请求参数"
//@Success 200
//@Router /signUp [post]
*/

func SignUpHandler(c *gin.Context) {
	//	获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//	请求参数有误,直接返回相应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//	判断err 是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} //这里有个小报错点:如果引入的包不是v10而是v8就会报错
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//	fmt.Printf(p.Password,p.Username,p.RePassword,p.Password)

	//业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//	返回响应

	//userID := snowflake.GenID()
	////    构造一个user实例
	//user := &models.User{
	//	UserID:   userID,
	//	Username: p.Username,
	//	Password: p.Password,
	//}
	//mysql.InsertUser(user)
	ResponseSuccess(c, nil)
}

//test
func Sign(c *gin.Context) {
	user := &models.User{
		UserID:   1223,
		Username: "sad",
		Password: "asdas",
	}
	mysql.InsertUser(user)
	ResponseSuccess(c, user)
	//mysql.SelectUserExist()
}

//登录
func LoginHandler(c *gin.Context) {
	//	获取参数以及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//	请求参数有误,直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		//	判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} //去掉报错信息的结构体字段,只保留后一个字段
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//	业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
