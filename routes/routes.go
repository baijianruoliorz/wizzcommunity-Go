package routes

import (
	"github.com/gin-contrib/pprof"
	"net/http"
	"wizzcommunity/controller"
	"wizzcommunity/logger"
	"wizzcommunity/middlewares"

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
	//cors
	v1.Use(middlewares.CORSMiddleware())
	//注册
	v1.POST("signUp", controller.SignUpHandler)
	//登录
	v1.POST("login", controller.LoginHandler)

	v1.GET("sign", controller.Sign)

	v1.GET("/post2", controller.GetPostListHandler2)
	v1.GET("/posts", controller.GetPostListHandler)
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	// 应用JWT认证中间件
	{
		v1.POST("/post", controller.CreatePostHandler)

		//  投票 按热度排序
		v1.POST("/vote", controller.PostVoteController)
	}
	//注册pprof相关路由
	pprof.Register(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"meg": "404",
		})
	})

	return r
}
