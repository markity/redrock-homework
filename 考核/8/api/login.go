package api

import (
	"github.com/gin-gonic/gin"
)

const lazyUsername = "admin"
const lazyPassword = "password"
const lazyAuthcode = "lazy-me"

func Login(ctx *gin.Context) {
	username, ok := ctx.GetPostForm("username")
	if !ok {
		ctx.String(200, "错误的用户名或密码")
		return
	}

	password, ok := ctx.GetPostForm("password")
	if !ok {
		ctx.String(200, "错误的用户名或密码")
		return
	}

	if username != lazyUsername || password != lazyPassword {
		ctx.String(200, "错误的用户名或密码")
		return
	}

	ctx.SetCookie("authcode", lazyAuthcode, 7200, "/", "localhost", false, false)
}
