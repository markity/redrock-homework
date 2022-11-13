package main

import (
	"lv1/handlers"
	"lv1/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化用户表
	utils.InitUserAuthMap()

	r := gin.Default()

	// 设置模板路径
	r.LoadHTMLGlob("templates/*")

	// 注册路由
	r.GET("/", handlers.Root)
	r.GET("/user", handlers.User)
	r.GET("/user/login", handlers.LoginGet)
	r.POST("/user/login", handlers.LoginPost)
	r.GET("/user/register", handlers.RegisterGet)
	r.POST("/user/register", handlers.RegisterPost)
	r.GET("/user/logout", handlers.Logout)

	r.Run()
}
