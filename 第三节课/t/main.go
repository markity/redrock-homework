package main

import (
	"sync"
	"time"
)

func main() {
	m := sync.Mutex{}

	m.Lock()

	go func() {
		println("进入1")
		time.Sleep(1 * time.Second)
		m.Lock()
		println(1)
		m.Unlock()
	}()

	go func() {
		println("进入2")
		time.Sleep(2 * time.Second)
		m.Lock()
		println(2)
		m.Unlock()
	}()

	go func() {
		println("进入3")
		time.Sleep(3 * time.Second)
		m.Lock()
		println(3)
		m.Unlock()
	}()

	go func() {
		println("进入4")
		time.Sleep(4 * time.Second)
		m.Lock()
		println(4)
		m.Unlock()
	}()

	go func() {
		println("进入5")
		time.Sleep(5 * time.Second)
		m.Lock()
		println(5)
		m.Unlock()
	}()

	time.Sleep(10 * time.Second)
	m.Unlock()

	time.Sleep(time.Second * 30)

}
