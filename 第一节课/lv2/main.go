package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b int
	var oper string
	n, _ := fmt.Scanf("%v", &a)
	if n != 1 {
		log.Fatalln("invaild")
	}
	n, _ = fmt.Scanf("%v", &oper)
	if n != 1 {
		log.Fatalln("invaild")
	}
	n, _ = fmt.Scanf("%v", &b)
	if n != 1 {
		log.Fatalln("invaild")
	}
	switch oper {
	case "+":
		println(a + b)
	case "-":
		println(a - b)
	case "*":
		println(a * b)
	case "/":
		println(a / b)
	default:
		println("无效的指令")
	}
}
