package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	c    chan struct{}
	flag bool
}

func (mm *MyMutex) Lock() {
	if !mm.flag {
		mm.flag = true
		mm.c = make(chan struct{}, 1)
	}
	mm.c <- struct{}{}
}

func (mm *MyMutex) Unlock() {
	<-mm.c
}

var x int64
var wg sync.WaitGroup
var mu MyMutex

func add() {
	for i := 0; i < 50000; i++ {
		mu.Lock()
		x = x + 1
		mu.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
