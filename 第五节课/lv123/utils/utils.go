package utils

import (
	"crypto/md5"
	"encoding/json"
	"log"
	"os"
	"sync"
	"unicode"
)

const FileName = "user.data"

var UserAuthMap map[string]string
var Mu sync.Mutex

// 将字符串MD5加密, 返回字符串
func ToMD5(a string) string {
	k := md5.New()
	k.Write([]byte(a))
	return string(k.Sum(nil))
}

// 检查用户名和密码在格式上是否合法
func CheckUnamePwdVaild(uname string, pwd string) bool {
	// 检查长度, uname长度为5~12, pwd长度为6~18
	if l1, l2 := len(uname), len(pwd); !(4 < l1 && l1 < 12) || !(5 < l2 && l2 < 18) {
		return false
	}

	// 用户名密码必须由字母数字组成
	for _, v := range uname {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) {
			return false
		}
	}
	for _, v := range pwd {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) {
			return false
		}
	}

	return true
}

// 初始化, UserAuathMap
func InitUserAuthMap() {
	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatalf("failed to ReadFile: %v\n", err)
	}

	if err := json.Unmarshal(data, &UserAuthMap); err != nil {
		log.Fatalf("failed to unmarshal json data: %v\n", err)
	}
}

// 刷新文件
func FlushFile() (e error) {
	Mu.Lock()
	defer Mu.Unlock()
	f, err := os.OpenFile(FileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		e = err
		return
	}

	// 关闭文件
	defer func() {
		if e == nil {
			e = f.Close()
		}
	}()

	b, err := json.Marshal(UserAuthMap)
	if err != nil {
		e = err
		return
	}

	_, err = f.Write(b)
	if err != nil {
		e = err
		return
	}

	e = nil
	return
}
