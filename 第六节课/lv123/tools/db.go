// include some useful tools and some program configs

package tools

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// global DataBase variable, be initialized in init() function
var db *sql.DB

// a cache to store security data
var securityQuestionsDataCache []string

// exec when package tools being imported, offering a global DataBase object
func init() {
	var err error
	db, err = sql.Open("mysql", dns)
	// fatal error, exit now
	if err != nil {
		log.Fatalf("failed to sql.Open: %v\n", err)
	}
	tryLoopPrepareSecuritryQuestionsCache()
}

// get *DB handler
func GetDBHandler() *sql.DB {
	return db
}

type User struct {
	ID        int
	Username  string
	AuthCode  string
	CreatedAt string
}

// get user info through username field
// if get none, return nil, true
// if get one, return &user, true
// otherwise return nil, false
func TryGetUserByUsername(username string) (*User, error) {
	var user User
	rows, err := db.Query("SELECT id, username, password_md5, created_at FROM user WHERE username=?", username)
	if err != nil {
		println(err)
		return nil, err
	}
	var data []byte
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &data, &user.CreatedAt); err != nil {
			if err != nil {
				println(err)
			}
			return nil, err
		}
		user.AuthCode = string(data)
		return &user, nil
	} else {
		return nil, nil
	}
}

// when project runs, load database security questions into the cache
// and it wound't change all the time
func tryLoopPrepareSecuritryQuestionsCache() {
	for {
		securityQuestionsDataCache = make([]string, 0)
		rows, err := db.Query("SELECT question FROM security ORDER BY id ASC")
		if err != nil {
			log.Printf("failed to tryLoopPrepareSecuritryQuestionsCache, retrying: %v\n", err)
			time.Sleep(time.Second * 3)
			continue
		}

		var tmp string
		retryflag := false
		for rows.Next() {
			if rows.Scan(&tmp) != nil {
				retryflag = true
				break
			}
			securityQuestionsDataCache = append(securityQuestionsDataCache, tmp)
		}

		if retryflag {
			continue
		} else {
			break
		}
	}
}

func Num2SecurityQuestion(i int) (string, bool) {
	if i > len(securityQuestionsDataCache)-1 || i < 0 {
		return "", false
	}
	return securityQuestionsDataCache[i], true
}

func GetSecurityQuestionsLen() int {
	return len(securityQuestionsDataCache)
}

// get the cache
func GetSecurityQuestionsDataCache() []string {
	return securityQuestionsDataCache
}
