package api

import "github.com/gin-gonic/gin"

func AuthMiddleware(ctx *gin.Context) {
	authcode, err := ctx.Cookie("authcode")
	if err != nil {
		ctx.String(200, "验证失败")
		ctx.Abort()
		return
	}

	if authcode != lazyAuthcode {
		ctx.String(200, "验证失败")
		ctx.Abort()
		return
	}

	ctx.Next()
}
