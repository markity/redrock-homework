// provide a auth middleware
// user not logined in will be redirected to a specified addr

package middleware

import (
	"database/sql"
	"log"
	"net/http"
	"project6/tools"

	"github.com/gin-gonic/gin"
)

// if a user does not logined in, redirect to login page
func GetMiddlewareMustLoginedIn(db *sql.DB, destAddr string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		println("fuck")
		username, err1 := ctx.Cookie("username")

		// authcode is not user's password, is is the md5 digest of user's password
		// database do not stroe user's password directly
		authcode, err2 := ctx.Cookie("authcode")

		// either of the cookie not exist, just redirect
		if err1 != nil || err2 != nil {
			println("re1" + destAddr)
			ctx.Redirect(http.StatusTemporaryRedirect, destAddr)
			ctx.Abort()
			return
		}

		// check the legality of username and authcode, reduce the pressure of the database
		if !tools.CheckUsernameValid(username) || !tools.CheckAuthcodeValid(authcode) {
			println("re2")
			ctx.Redirect(http.StatusTemporaryRedirect, destAddr)
			ctx.Abort()

			// do not pass to the next middleware, just exit
			return
		}

		// attempt specified times
		user, err := tools.TryGetUserByUsername(username)
		if err != nil {
			log.Panicf("failed to Query: %v\n", err)
		}

		// user authcode invalid, redirect
		if user == nil || (user.AuthCode != authcode) {
			println("re3")
			ctx.Redirect(http.StatusTemporaryRedirect, destAddr)
			ctx.Abort()
			return
		}

		println("aqq")
		// ok, logined in, pass it to next middleware
		// and need to update cookie expire time to keep seesion active
		ctx.SetCookie("username", username, 3600, "/", "localhost", false, false)
		ctx.SetCookie("authcode", user.AuthCode, 3600, "/", "localhost", false, false)
		ctx.Set("user", user)
		ctx.Next()
	}
}

func GetMiddlewareMustNotLoginedIn(db *sql.DB, destAddr string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		username, err1 := ctx.Cookie("username")

		// authcode is not user's password, is is the md5 digest of user's password
		// database do not stroe user's password directly
		authcode, err2 := ctx.Cookie("authcode")

		// either of the cookie not exist, mean not logined in
		if err1 != nil || err2 != nil {
			ctx.Next()
			return
		}

		// check the legality of username and authcode, reduce the pressure of the database
		if !tools.CheckUsernameValid(username) || !tools.CheckAuthcodeValid(authcode) {
			ctx.Next()
			return
		}

		// attempt specified times
		user, err := tools.TryGetUserByUsername(username)
		if err != nil {
			log.Panicf("failed to Query: %v\n", err)
		}

		// user authcode invalid, redirect
		if user == nil || user.AuthCode != authcode {
			ctx.Next()
			return
		}

		// logined in, redirect to destAddr
		ctx.SetCookie("username", "", -1, "/", "localhost", false, false)
		ctx.SetCookie("authcode", "", -1, "/", "localhost", false, false)
		ctx.Redirect(http.StatusMovedPermanently, destAddr)
	}
}
