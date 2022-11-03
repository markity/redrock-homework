package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cancel1, jump1 := NewDailyScheduler(8, 30, 0, func() {
		println("八点三十零秒到了")
	}).Run()

	// 跳过一次闹铃, 在下下次八点30会调度
	jump1()

	cancel2, _ := NewDailyScheduler(9, 40, 0, func() {
		println("九点四十到了零秒到了")
	}).Run()

	// TickerScheduler也很容易实现jump逻辑, 这里不写了
	cancel3 := NewTickerScheduler(time.Second*30, func() {
		println("30s过去了")
	}).Run()

	cancel4 := NewTickerScheduler(time.Second*10, func() {
		println("10s过去了")
	}).Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// 一直阻塞, 直到外部信号通知结束程序, 此时关闭所有定时器
	<-c
	cancel1()
	cancel2()
	cancel3()
	cancel4()
	println("程序结束")
}
