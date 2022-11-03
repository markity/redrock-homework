package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

var UserData map[string]*UserInfo

const FileName = "user.data"

// 获得文件内容
func LoadFile() (by []byte, er error) {
	f, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	// 关闭文件
	defer func() {
		if er == nil {
			er = f.Close()
		}
	}()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		by = nil
		er = err
		return
	}

	by = b
	er = nil
	return
}

func FlushFile() (e error) {
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

	udj := UserDataJson{User: UserData}

	b, err := json.Marshal(&udj)
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

func GetMD5(s string) string {
	p := md5.New()
	p.Write([]byte(s))
	return string(p.Sum(nil))
}

// 用户信息
type UserInfo struct {
	UserName string `json:"-"`
	NickName string `json:"nick_name"`
	PassMd5  string `json:"pass_md5"`
}

func (ui *UserInfo) String() string {
	return fmt.Sprintf("用户名: %v, 显示名: %v", ui.UserName, ui.NickName)
}

// 用于user.data的Json解析
type UserDataJson struct {
	User map[string](*UserInfo) `json:"user"`
}

// 统一定义一些用户提示信息

const Msg1 = `请输入你的操作:
1. 注册
2. 登录
3. 退出
`

const Msg2 = `请输入用户名(要求只允许使用数字下划线和字母, 长度为5-12个字符):
`

const Msg3 = `请输入密码(长度6-20个字节):
`

const Msg4 = `请重复密码:
`

const Msg5 = `请输入你想要展示给外界的名字(4-16个字节, 且不允许包含空格或制表符):
`

const Msg6 = `不允许输入空行, 请重新输入:
`

const Msg7 = `用户名已存在, 请重新输入:
`

const Msg8 = `用户名不合法, 请重新输入:
`

const Msg9 = `不合法的密码, 请重新输入:
`

const Msg10 = `与前文密码不符, 请重新输入:
`

const Msg11 = `不合法的显示名, 请重新输入:
`

const Msg12 = `请输入用户名:
`

const Msg13 = `请输入密码:
`

const Msg14 = `错误的用户名或密码
`

const Msg15 = `恭喜你成功登录, 你可以进行以下操作:
1. 查看账户信息
2. 查看可执行的操作
3. 修改显示名称
4. 修改密码
5. 退出登录
6. 退出程序
`

const Msg16 = `未知的指令, 请重新输入
`

const Msg17 = `你可以继续操作:
1. 查看账户信息
2. 查看可执行的操作
3. 修改显示名称
4. 修改密码
5. 退出登录
6. 退出程序
`

const Msg18 = `请输入下个操作:
`

const Msg19 = `1. 查看账户信息
2. 查看可执行的操作
3. 修改显示名称
4. 修改密码
5. 退出登录
6. 退出程序
`

// 要求只允许使用数字下划线和字母, 长度为5-12个字符
// 第一返回值是是否合法, 第二个返回值是是否已经存在
func checkNewUsernameVaild(s string) (bool, bool) {
	cnt := 0
	// 要求
	for _, v := range s {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) && v != '_' {
			return false, false
		}
		cnt++
	}

	if !(cnt >= 5 && cnt <= 12) {
		return false, false
	}

	if _, ok := UserData[s]; ok {
		return true, true
	}

	return true, false
}

// 检查密码是否合法, 要求密码大小6到20个字节
func checkPasswordVaild(s string) bool {
	if l := len(s); l < 6 && l > 20 {
		return false
	}
	return true
}

// 检查nickname是否合法, 不允许包含空格或制表符
func checkNickNameVaild(s string) bool {
	if strings.Contains(s, " ") || strings.Contains(s, "\t") {
		return false
	}
	if l := len(s); l < 4 || l > 16 {
		return false
	}
	return true
}

// 读取一个line
func scanNewLine(sc *bufio.Scanner) (string, error) {
	if sc.Scan(); sc.Err() != nil {
		return "", sc.Err()
	}
	return sc.Text(), nil
}

func doRegister(sc *bufio.Scanner) {
	var username string
	var password string
	var nickName string
	print(Msg2)
	// 获取username
	for {
		un, err := scanNewLine(sc)
		if err != nil {
			log.Fatalf("failed to scan line: %v\n", err)
		}

		vaild, existed := checkNewUsernameVaild(un)
		if !vaild {
			print(Msg8)
			continue
		}
		if existed {
			print(Msg7)
			continue
		}

		username = un
		break
	}

	print(Msg3)
	// 获取password
	for {
		pwd, err := scanNewLine(sc)
		if err != nil {
			log.Fatalf("failed to scan line: %v\n", err)
		}
		if pwd == "" {
			print(Msg6)
			continue
		}

		vaild := checkPasswordVaild(pwd)
		if !vaild {
			print(Msg9)
		}

		password = pwd
		break
	}

	print(Msg4)
	// 获取passwordConfirm
	for {
		pwd, err := scanNewLine(sc)
		if err != nil {
			log.Fatalf("failed to scan line: %v\n", err)
		}
		if pwd == "" {
			print(Msg6)
			continue
		}

		if pwd != password {
			print(Msg10)
			continue
		}
		break
	}

	print(Msg5)
	// 获取nickName
	for {
		nick, err := scanNewLine(sc)
		if err != nil {
			log.Fatalf("failed to scan line: %v\n", err)
		}
		if nick == "" {
			print(Msg6)
			continue
		}

		if !checkNickNameVaild(nick) {
			print(Msg11)
			continue
		}

		nickName = nick
		break
	}

	// 注册
	ui := UserInfo{
		NickName: nickName,
		PassMd5:  GetMD5(password),
	}
	UserData[username] = &ui
	err := FlushFile()
	if err != nil {
		log.Fatalf("failed to flush file: %v\n", err)
	}
}

func doLogin(sc *bufio.Scanner) (bool, *UserInfo) {
	var username string
	var password string
	print(Msg12)
	// 获取username
	for {
		un, err := scanNewLine(sc)
		if err != nil {
			log.Fatalf("failed to scan line: %v\n", err)
		}

		// 不允许空行
		if un == "" {
			print(Msg6)
			continue
		}

		username = un
		break
	}

	print(Msg13)
	// 获取password
	for {
		pwd, err := scanNewLine(sc)
		if err != nil {
			log.Fatalf("failed to scan line: %v\n", err)
		}

		if pwd == "" {
			print(Msg6)
			continue
		}

		vaild := checkPasswordVaild(pwd)
		if !vaild {
			print(Msg14)
			return false, nil
		}

		password = pwd
		break
	}

	ui, ok := UserData[username]
	if !ok {
		print(Msg14)
		return false, nil
	}

	if (GetMD5(password)) != ui.PassMd5 {
		print(Msg14)
		return false, nil
	}

	ui.UserName = username

	return true, ui
}

// 返回值表示是否退出程序
func doAfterLogin(sc *bufio.Scanner, ui *UserInfo) bool {
	firstEnter := true
	print(Msg15)
	for {
		if !firstEnter {
			print(Msg18)
		}
		firstEnter = false
		option, err := scanNewLine(sc)
		if err != nil {
			log.Fatalf("failed to scanNewLine")
		}
		switch option {
		// 查看信息
		case "1":
			fmt.Printf("你好, 用户%v, 你的nickname为: %v\n", ui.UserName, ui.NickName)
			continue
		// 提示信息
		case "2":
			print(Msg19)
			continue
		// 修改显示名称
		case "3":
			print(Msg5)

			// 获取nickName
			for {
				nick, err := scanNewLine(sc)
				if err != nil {
					log.Fatalf("failed to scan line: %v\n", err)
				}
				if nick == "" {
					print(Msg6)
					continue
				}

				if !checkNickNameVaild(nick) {
					print(Msg11)
					continue
				}

				ui.NickName = nick
				break
			}

			err = FlushFile()
			if err != nil {
				log.Fatalf("failed to flush file: %v\n", err)
			}

			continue
		// 修改密码
		case "4":
			print(Msg3)
			var password string
			// 获取password
			for {
				pwd, err := scanNewLine(sc)
				if err != nil {
					log.Fatalf("failed to scan line: %v\n", err)
				}
				if pwd == "" {
					print(Msg6)
					continue
				}

				vaild := checkPasswordVaild(pwd)
				if !vaild {
					print(Msg9)
				}

				password = pwd
				break
			}

			print(Msg4)
			// 获取passwordConfirm
			for {
				pwd, err := scanNewLine(sc)
				if err != nil {
					log.Fatalf("failed to scan line: %v\n", err)
				}
				if pwd == "" {
					print(Msg6)
					continue
				}

				if pwd != password {
					print(Msg10)
					continue
				}
				break
			}

			ui.PassMd5 = GetMD5(password)

			err := FlushFile()
			if err != nil {
				log.Fatalf("failed to flush file: %v\n", err)
			}
			continue
		// 退出登录
		case "5":
			return false
		// 退出程序
		case "6":
			return true
		default:
			print(Msg16)
			continue
		}
	}
}

func main() {
	// 加载文件
	data, err := LoadFile()
	if err != nil {
		log.Fatalf("failed to load file: %v\n", err)
	}

	// 解析json
	udj := UserDataJson{}
	if err = json.Unmarshal(data, &udj); err != nil {
		log.Fatalf("failed to unmarshal json data: %v\n", err)
	}

	// 更新全局用户变量
	UserData = udj.User

	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(Msg1)
		option, err := scanNewLine(sc)
		if err != nil {
			log.Fatalf("failed to scan new line: %v\n", err)
		}
		switch option {
		case "1":
			// 返回即为注册完毕
			doRegister(sc)
			continue
		case "2":
			ok, ui := doLogin(sc)
			if !ok {
				continue
			}

			// 成功登录
			exit := doAfterLogin(sc, ui)
			if exit {
				return
			} else {
				continue
			}
		case "3":
			return
		default:
			print(Msg16)
		}
	}
}
