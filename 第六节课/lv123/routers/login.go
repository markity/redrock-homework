package routers

import (
	"log"
	"net/http"
	"project6/tools"

	"github.com/gin-gonic/gin"
)

// be warpped by unlogin auth middleware
func loginGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.tmpl", nil)
}

// be warpped by login auth middleware
/* postform structure
username
password
*/
func loginPost(ctx *gin.Context) {
	username, ok1 := ctx.GetPostForm("username")
	password, ok2 := ctx.GetPostForm("password")

	if !ok1 || !ok2 ||
		!tools.CheckUsernameValid(username) || !tools.CheckPasswordValid(password) {
		ctx.String(http.StatusBadRequest, "invalid post data")
		return
	}

	// to notify the latter control flow, attempts success
	user, err := tools.TryGetUserByUsername(username)
	if err != nil {
		log.Panicf("failed to Query: %v\n", err)
	}

	if user == nil || (user.AuthCode != tools.GetMD5Digest(password)) {
		ctx.HTML(http.StatusBadRequest, "wrong_password.tmpl", nil)
		return
	}

	// ok, set cookie
	ctx.SetCookie("username", username, 3600, "/", "localhost", false, false)
	ctx.SetCookie("authcode", user.AuthCode, 3600, "/", "localhost", false, false)
	ctx.HTML(http.StatusOK, "login_success.tmpl", nil)
}
