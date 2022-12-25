package main

import (
	"depo/api"
	"depo/service"

	"github.com/gin-gonic/gin"
)

func main() {
	service.MustPrepareTables()
	engine := gin.Default()
	api.InitGroup(engine)

	engine.Run("127.0.0.1:8000")
}
