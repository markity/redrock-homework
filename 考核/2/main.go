package main

import "fmt"

// 原因: 根本没有执行defer语句, 就return出去了

func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()

	if a {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()
}
