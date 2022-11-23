package routers

import (
	"log"
	"net/http"
	"project6/tools"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// be wrapped by unlogin auth middleware
func registerGet(ctx *gin.Context) {
	cache := tools.GetSecurityQuestionsDataCache()
	for k, v := range cache {
		println(k, v)
	}
	ctx.HTML(http.StatusOK, "register.tmpl", gin.H{
		"security_questions": cache,
	})
}

// postform structure
/*
username
password
confirm_password
security_choice
security_answer
*/
func registerPost(ctx *gin.Context) {
	// get postform data
	username, ok1 := ctx.GetPostForm("username")
	password, ok2 := ctx.GetPostForm("password")
	confirmPassword, ok3 := ctx.GetPostForm("confirm_password")
	securityChoice, ok4 := ctx.GetPostForm("security_choice")
	securityAnswer, ok5 := ctx.GetPostForm("security_answer")
	var securityChoiceNum int

	// basic check, to reduce database pressure
	var err error
	if securityChoiceNum, err = strconv.Atoi(securityChoice); !ok1 || !ok2 || !ok3 || !ok4 || !ok5 ||
		!tools.CheckUsernameValid(username) ||
		!tools.CheckPasswordValid(password) ||
		!tools.CheckSecurityAnswerValid(securityAnswer) ||
		password != confirmPassword ||
		err != nil ||
		!tools.CheckSecurityAnswerValid(securityAnswer) {

		ctx.HTML(http.StatusBadRequest, "wrong_register_input.tmpl", nil)
		return
	}
	_, ok := tools.Num2SecurityQuestion(securityChoiceNum)
	if !ok {
		ctx.HTML(http.StatusBadRequest, "wrong_register_input.tmpl", nil)
		return
	}

	// golang format time 2006-01-02 15:04:05
	rows, err := tools.GetDBHandler().Query("SELECT username, password_md5, created_at FROM user WHERE username = ?", username)
	if err != nil {
		log.Panicf("failed to Query: %v\n", err)
	}

	// the username had been used
	if rows.Next() {
		ctx.HTML(http.StatusOK, "username_in_use.tmpl", nil)
		return
	}

	res, err := tools.GetDBHandler().Exec("INSERT INTO user(username, password_md5, created_at) VALUES (?, ?, ?)", username, tools.GetMD5Digest(password),
		time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Panicf("failed to Exec: %v\n", err)
		return
	}
	// ignore this error, mysql supports thi
	lastInserted, _ := res.LastInsertId()

	_, err = tools.GetDBHandler().Exec("INSERT INTO security_user(user_id, security_id, answer) VALUES(?, ?, ?)", lastInserted, securityChoiceNum+1, securityAnswer)
	if err != nil {
		log.Panicf("failed to Exec: %v\n", err)
		return
	}

	// success, so redirect
	ctx.HTML(http.StatusOK, "register_success.tmpl", nil)
}
