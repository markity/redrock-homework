package handlers

import (
	"lv123/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	// cookie保存的用户名
	username, err1 := c.Cookie("username")

	// cookie保存用户的认证码, 即为password简单加密
	authcode, err2 := c.Cookie("authcode")

	// 防止竞态, 读出错误的数据
	utils.Mu.Lock()
	uapAuthcode, ok := utils.UserAuthMap[username]
	defer utils.Mu.Unlock()

	// 如果鉴权失败, 即没有登录
	if err1 != nil || err2 != nil || !ok || uapAuthcode != authcode {
		c.Redirect(http.StatusTemporaryRedirect, "/user/login")
		return
	}

	// 鉴权成功时返回用户信息
	c.HTML(http.StatusOK, "user_page.tmpl", gin.H{
		"username": username,
	})
}
