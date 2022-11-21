package main

import (
	"project6/routers"
	"project6/tools"

	"github.com/gin-gonic/gin"
)

func main() {
	// register routers
	db := tools.GetDBHandler()
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	group := engine.Group("/")
	routers.InitGroup(group, db)

	// run server
	engine.Run()
}
