package handlers

import (
	"lv123/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()
	username, err1 := c.Cookie("username")

	authcode, err2 := c.Cookie("authcode")

	uapAuthcode, ok := utils.UserAuthMap[username]

	// 如果没有登录
	if err1 != nil || err2 != nil || !ok || uapAuthcode != authcode {
		c.String(http.StatusOK, "you have not logined yet")
		return
	}

	// 删除cookie
	c.SetCookie("username", "", -1, "/", "localhost", false, false)
	c.SetCookie("authcode", "", -1, "/", "localhost", false, false)

	c.Redirect(http.StatusTemporaryRedirect, "/user/login")
}
