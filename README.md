Vision: Do something for wizz studio

Q:为什么还做Go 版本/kk

A:如果做小程序可以适当阉割一些功能,网页版用SpringBoot的后台 ~~并不是gin太蔡了~~

简单的上手:

1.air使用简单: git clone下代码后 下载好需要的依赖包 只需在控制台输入air即可启动本项目

关于air:https://github.com/cosmtrek/air

2.swag init 可生成相应API 不过不如Java的方便.



update:10.22 

忙里偷闲的复习之余把小程序的go后台开发完了,虽然偷懒没有怎么写swag注释(太麻烦)

目前仅仅注册了两个方法和若干结构体的swagger:http://112.126.78.122:8083/swagger/index.html

复习两小时奖励自己写四十分钟代码?

Java的主后台还是留给期中之后吧~ 

欢迎有想法的大佬git clone二次开发吊打我orz 

wizzCommunity built by Go             --baijianruoliorz

update 2020.10.25:
增加了一个gormDemo包,因为这个项目的orm用的是sqlx(个人比较习惯)

但是考虑到gorm非常方便,所以写了两个小demo 几乎囊括了日常的gorm使用

以及重新编写了DockerFile 它现在变得更符合这个项目了(~~之前偷懒直接复制其他的项目~~)

因为我觉得这样很coooooooool!
