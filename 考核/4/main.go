package main

import (
	"fmt"
	"time"
)

func main() {
	// 用channel协调任务的顺序
	do1 := make(chan struct{})
	ok1 := make(chan struct{})
	do2 := make(chan struct{})
	ok2 := make(chan struct{})
	do3 := make(chan struct{})
	ok3 := make(chan struct{})
	go Work("goroutine1", do1, ok1)
	go Work("goroutine2", do2, ok2)
	go Work("goroutine3", do3, ok3)

	do1 <- struct{}{}
	<-ok1

	do2 <- struct{}{}
	<-ok2

	do3 <- struct{}{}
	<-ok3

	fmt.Println("successful")
}

func Work(workName string, do chan struct{}, ok chan struct{}) {
	<-do

	fmt.Println(workName)
	time.Sleep(time.Second) // 模拟业务

	ok <- struct{}{}
}
