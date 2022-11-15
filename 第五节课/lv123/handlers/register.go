package handlers

import (
	"log"
	"lv123/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGet(c *gin.Context) {
	username, err1 := c.Cookie("username")
	authcode, err2 := c.Cookie("authcode")

	utils.Mu.Lock()
	uapAuthcode, ok := utils.UserAuthMap[username]
	utils.Mu.Unlock()

	// 鉴权失败才返回注册模板
	if err1 != nil || err2 != nil || !ok || uapAuthcode != authcode {
		c.HTML(http.StatusOK, "register_page.tmpl", nil)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/user")
}

func RegisterPost(c *gin.Context) {
	username, err1 := c.Cookie("username")
	authcode, err2 := c.Cookie("authcode")

	utils.Mu.Lock()
	uapAuthcode, ok := utils.UserAuthMap[username]
	utils.Mu.Unlock()

	// 判断特殊情况, 已经登录的人再来发表单
	if !(err1 != nil || err2 != nil || !ok || uapAuthcode != authcode) {
		c.String(http.StatusOK, "you hava already authed! please logout first")
		return
	}

	// 检查表单内容是否正确
	usernameArray := c.PostFormArray("username")
	passwordArray := c.PostFormArray("password")
	if len(usernameArray) != 1 || len(passwordArray) != 1 {
		c.String(http.StatusOK, "wrong login paraments")
		utils.Mu.Unlock()
		return
	}
	uname := usernameArray[0]
	pwd := passwordArray[0]

	// 这里是一个先读, 如果没有则写的情况, 应一直锁住, 防止幻读
	utils.Mu.Lock()
	_, existed := utils.UserAuthMap[uname]
	if existed {
		c.String(http.StatusOK, "the username existed already")
		utils.Mu.Unlock()
		return
	}

	if !utils.CheckUnamePwdVaild(uname, pwd) {
		c.String(http.StatusOK, "the format of the password or username is wrong")
		utils.Mu.Unlock()
		return
	}

	acode := utils.ToMD5(pwd)
	utils.UserAuthMap[uname] = acode
	utils.Mu.Unlock()
	if err := utils.FlushFile(); err != nil {
		log.Printf("failed to sync file: %v\n", err)
	}

	// 设置cookie
	c.SetCookie("username", uname, 3600, "/", "localhost", false, false)
	c.SetCookie("authcode", acode, 3600, "/", "localhost", false, false)

	// 请求方法改变,
	c.Redirect(http.StatusMovedPermanently, "/user")
}
