package api

import "github.com/gin-gonic/gin"

func NewCargo(ctx *gin.Context) {
	depoName, ok := ctx.GetPostForm("depo_name")
	if !ok {
		ctx.String(200, "错误的输入")
	}
}
