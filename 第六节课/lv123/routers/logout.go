package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// be warpped by auth middleware
func logoutGet(ctx *gin.Context) {
	ctx.SetCookie("username", "", -1, "/", "localhost", false, false)
	ctx.SetCookie("authcode", "", -1, "/", "localhost", false, false)
	ctx.HTML(http.StatusOK, "logout_success.tmpl", nil)
}
