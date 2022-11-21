// a script to create tables and prepare security questions
// should be run before run ../main.go project
// if panic when running this, create database named project6 and try again

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dns = "root:mark2004119@/project6"
var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", dns)
	// fatal error, exit now
	if err != nil {
		log.Fatalf("failed to sql.Open: %v\n", err)
	}

	// do migrate tables
	DropAllTables()
	AutoMigrateTables()

	// prepare security questions
	var securityQuestions = []string{
		"你母亲出生的年份",
		"你宠物狗的名字",
		"你的小学母校全名",
	}
	for _, v := range securityQuestions {
		_, err := db.Exec("INSERT INTO security(question) VALUES (?)", v)
		if err != nil {
			log.Fatalf("failed to insert into security table: %v", err)
		}
	}
}

// a tool to try to create a table
// if failed, print log and retry
func tryLoopCreateTables(s string) {
	for {
		_, err := db.Exec(s)
		if err != nil {
			log.Printf("failed to tryLoopCreateTables, retrying: %v\n", err)
			log.Printf("%s", s)
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}
}

// a tool to try to drop a table
// if failed, print log and retry
func tryLoopDropTables(s string) {
	for {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v", s))
		if err != nil {
			log.Printf("failed to tryLoopDropTables, retrying: %v\n", err)
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}
}

// to create tables. if errors occure, always retry
// if table exists, it does nothing
func AutoMigrateTables() {
	// table user
	tryLoopCreateTables(migrateTableUser)
	// table scurity_user
	tryLoopCreateTables(migrateTableSecurityUser)
	// table security
	tryLoopCreateTables(migrateTableSecurity)
	// table comment
	tryLoopCreateTables(migreateTableComment)
}

// to drop all tables
func DropAllTables() {
	// table user
	tryLoopDropTables("user")
	//table security_user
	tryLoopDropTables("security_user")
	// table security
	tryLoopDropTables("security")
	// table comment
	tryLoopDropTables("comment")
}
