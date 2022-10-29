package main

import (
	"fmt"
	"sync"
)

var count int = 1

var c chan struct{} = make(chan struct{})

func PrintJ(wg *sync.WaitGroup) {
	for {
		// 奇数先打, 打了等待对方执行完毕
		fmt.Println(count)
		count++
		// 退出携程
		if count == 100 {
			c <- struct{}{}
			wg.Done()
			return
		}
		// 发送完毕后等待对方发送信号
		c <- struct{}{}
		<-c
	}
}

func PrintO(wg *sync.WaitGroup) {
	for {
		// 偶数后打, 收到信号之后向对方发送信号
		<-c
		fmt.Println(count)
		if count == 100 {
			wg.Done()
			return
		}
		c <- struct{}{}
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go PrintJ(&wg)
	go PrintO(&wg)
	wg.Wait()
}
