package routers

import (
	"log"
	"net/http"
	"project6/tools"

	"github.com/gin-gonic/gin"
)

// be wrapped by auth middleware
func changePWDGet(ctx *gin.Context) {
	user_, _ := ctx.Get("user")
	user := user_.(*tools.User)
	var question string
	err := tools.GetDBHandler().QueryRow(
		`SELECT security.question AS question FROM user, security_user, security WHERE
			 user.id=security_user.user_id AND security.id=security_user.security_id AND user.id=?`, user.ID).Scan(&question)
	if err != nil {
		log.Panicf("failed to query, retrying: %v\n", err)
	}
	ctx.HTML(http.StatusOK, "changepwd.tmpl", gin.H{
		"security_question": question,
	})
}

// be wrapped by auth middleware
/* postform structure
username
old_password
*/
func changePWDPost(ctx *gin.Context) {
	username, ok1 := ctx.GetPostForm("username")
	oldPassword, ok2 := ctx.GetPostForm("old_password")
	newPassword, ok3 := ctx.GetPostForm("new_password")
	confirmPassword, ok4 := ctx.GetPostForm("confirm_password")
	securityAnswer, ok5 := ctx.GetPostForm("security_answer")

	// basic check to reduce the pressure of database
	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 ||
		!tools.CheckUsernameValid(username) ||
		!tools.CheckPasswordValid(oldPassword) ||
		!tools.CheckPasswordValid(newPassword) ||
		!tools.CheckPasswordValid(confirmPassword) ||
		!tools.CheckSecurityAnswerValid(securityAnswer) ||
		newPassword != confirmPassword {
		ctx.HTML(http.StatusBadRequest, "wrong_changepwd_input.tmpl", nil)
		return
	}

	user_, _ := ctx.Get("user")
	user := user_.(*tools.User)
	if tools.GetMD5Digest(oldPassword) != user.AuthCode {
		ctx.HTML(http.StatusBadRequest, "wrong_changpwd_input.tmpl", nil)
	}

	var rightAnwer string
	// inner join three tables
	err := tools.GetDBHandler().QueryRow(
		`SELECT security_user.answer AS answer FROM user, security_user, security WHERE
			 user.id=security_user.user_id AND security.id=security_user.security_id AND user.id=?`, user.ID).Scan(&rightAnwer)
	if err != nil {
		log.Panicf("failed to query, retrying: %v\n", err)
	}

	_, err = tools.GetDBHandler().Exec("UPDATE user SET password_md5=? WHERE id=?", tools.GetMD5Digest(newPassword), user.ID)
	if err != nil {
		log.Panicf("failed to Update: %v\n", err)
	}

	ctx.SetCookie("username", "", -1, "/", "localhost", false, false)
	ctx.SetCookie("authcode", "", -1, "/", "localhost", false, false)
	ctx.HTML(http.StatusOK, "changepwd_success.tmpl", nil)
}
