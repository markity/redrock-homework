package api

import (
	"depo/dao"

	"github.com/gin-gonic/gin"
)

func NewDepository(ctx *gin.Context) {
	depoName, ok := ctx.GetPostForm("depo_name")
	if !ok {
		ctx.String(200, "错误的输入")
		return
	}

	_, err := dao.DB.Exec("INSERT INTO depository(name) VALUES (?)", depoName)
	if err != nil {
		ctx.String(200, "服务器错误")
		return
	}
}
