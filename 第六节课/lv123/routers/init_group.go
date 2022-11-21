package routers

import (
	"database/sql"
	"project6/middleware"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func InitGroup(g *gin.RouterGroup, db *sql.DB) {
	DB = db
	g.GET("/", middleware.GetMiddlewareMustLoginedIn(db, "/login"), mainPageGet)
	g.POST("/", middleware.GetMiddlewareMustLoginedIn(db, "/login"), mainPagePost)

	g.GET("/changepwd", middleware.GetMiddlewareMustLoginedIn(db, "/login"), changePWDGet)
	g.POST("/changepwd", middleware.GetMiddlewareMustLoginedIn(db, "/login"), changePWDPost)

	g.GET("/login", middleware.GetMiddlewareMustNotLoginedIn(db, "/"), loginGet)
	g.POST("/login", middleware.GetMiddlewareMustNotLoginedIn(db, "/"), loginPost)

	g.GET("/logout", middleware.GetMiddlewareMustLoginedIn(db, "/login"), logoutGet)

	g.GET("/register", middleware.GetMiddlewareMustNotLoginedIn(db, "/"), registerGet)
	g.POST("/register", middleware.GetMiddlewareMustNotLoginedIn(db, "/"), registerPost)
}
