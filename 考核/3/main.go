package main

import "fmt"

func main() {
	c := make(chan struct{})

	go func() {
		fmt.Println("出现")
		c <- struct{}{}
	}()

	<-c
}
