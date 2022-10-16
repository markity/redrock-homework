package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var user = map[string]string{
	"Markity": "It_Is_pass",
	"lwaix":   "pass_too",
}

func readLine(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	return reader.ReadString('\n')
}

func main() {
	username, err := readLine(os.Stdin)
	if err != nil {
		fmt.Printf("failed to readLine: %v\n", err)
		os.Exit(-1)
	}

	password, err := readLine(os.Stdin)
	if err != nil {
		fmt.Printf("failed to readLine: %v\n", err)
		os.Exit(-1)
	}

	// 去\n
	username = username[:len(username)-1]
	password = password[:len(password)-1]

	if user[username] == password {
		println("密码正确")
	} else {
		println("密码错误")
	}
}
