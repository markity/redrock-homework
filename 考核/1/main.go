package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// 问题: 字段名小写, 不能导出, json包扫不出来
type user struct {
	Username string
	Nickname string
	Sex      uint8
	Birthday time.Time
}

func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
