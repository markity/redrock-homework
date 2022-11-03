package main

import (
	"sync"
)

var count int = 1

var c chan struct{} = make(chan struct{})

func PrintJ(wg *sync.WaitGroup) {
	a := 1
	for {
		// 奇数先开始
		println(a)
		if a == 99 {
			c <- struct{}{}
			wg.Done()
			return
		}
		a += 2

		c <- struct{}{}

		<-c
	}
}

func PrintO(wg *sync.WaitGroup) {
	b := 2
	for {
		// 偶数后打, 收到信号之后向对方发送信号
		<-c
		println(b)
		if b == 100 {
			wg.Done()
			return
		}
		b += 2

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
