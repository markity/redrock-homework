package main

import (
	"fmt"
	"sync"
)

func main() {
	over := make(chan bool)
	go func() {
		// 分发任务
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int, wg *sync.WaitGroup) {
				fmt.Println(i)
				wg.Done()
			}(i, &wg)
		}

		// 等待全部任务完成
		wg.Wait()
		over <- true

	}()
	<-over
	fmt.Println("over!!!")
}
