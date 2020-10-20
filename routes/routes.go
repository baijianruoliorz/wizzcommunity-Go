package routes

import (
	"net/http"
	"wizzcommunity/controller"
	"wizzcommunity/logger"

	"github.com/gin-gonic/gin"

	//swagger需要的
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "wizzcommunity/docs"
)

func Setup() *gin.Engine {
	r := gin.New()

	//两个记录日志以及让日志recovery的方法
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello guy,I'm a developer who is exploitting this project.\n"+
			"You can find me in the github:https://github.com/baijianruoliorz.\n"+
			"I am a student from xidian university,China.\n"+
			"And i am looking for a full-time job for my internship experience.\n"+
			"How to touch me:QQ:1099462011.")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	//注册
	v1.POST("signUp", controller.SignUpHandler)
	//登录
	v1.POST("login", controller.LoginHandler)
	return r
}
