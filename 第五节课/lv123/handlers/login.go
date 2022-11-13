package handlers

import (
	"lv1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginGet(c *gin.Context) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()
	// cookie保存的用户名
	username, err1 := c.Cookie("username")

	// cookie保存用户的认证码, 即为password简单加密
	authcode, err2 := c.Cookie("authcode")

	uapAuthcode, ok := utils.UserAuthMap[username]

	// 鉴权失败才返回登录模板
	if err1 != nil || err2 != nil || !ok || uapAuthcode != authcode {
		c.HTML(http.StatusOK, "login_page.tmpl", nil)
	}

	// 鉴权成功, 已经登录, 跳转到/user
	c.Redirect(http.StatusTemporaryRedirect, "/user")
}

func LoginPost(c *gin.Context) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()
	username, err1 := c.Cookie("username")

	authcode, err2 := c.Cookie("authcode")

	uapAuthcode, ok := utils.UserAuthMap[username]

	// 判断特殊情况, 已经登录的人再来发表单
	if !(err1 != nil || err2 != nil || !ok || uapAuthcode != authcode) {
		c.String(http.StatusBadRequest, "you hava already authed! please logout first")
		return
	}

	// 检查表单内容是否正确
	usernameArray := c.PostFormArray("username")
	passwordArray := c.PostFormArray("authcode")
	if len(usernameArray) != 1 || len(passwordArray) != 1 {
		c.String(http.StatusUnauthorized, "wrong login paraments")
		return
	}
	uname := usernameArray[0]
	acode := utils.ToMD5(passwordArray[0])
	acode_, ok := utils.UserAuthMap[uname]
	if !ok || acode != acode_ {
		c.String(http.StatusUnauthorized, "faild to login, check your password again")
		return
	}

	// 设置cookie
	c.SetCookie("username", uname, 3600, "/", "localhost", false, false)
	c.SetCookie("authcode", acode_, 3600, "/", "localhost", false, false)

	c.Redirect(http.StatusMovedPermanently, "/user")
}
