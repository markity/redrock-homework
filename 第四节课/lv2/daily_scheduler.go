package main

import (
	"fmt"
	"time"
)

func getSoleTick(t time.Duration) chan struct{} {
	c := make(chan struct{})
	go func() {
		time.Sleep(t)
		c <- struct{}{}
		close(c)
	}()
	return c
}

type dailyScheduler struct {
	hour   int
	min    int
	second int
	f      func()
}

// 非阻塞运行
func (d dailyScheduler) Run() (func(), func()) {
	// 如果今天的时间已经过了规定时间, 那么等明天
	now := time.Now().Local()
	// temp假设如果没过今天, 那么应该回调的时间是多少
	temp := time.Date(now.Year(), now.Month(), now.Day(), d.hour, d.min, d.second, 0, time.Local)
	// 回调的真实时间
	var ringDate time.Time
	// 判断是否已经过了今天的时间
	if now.After(temp) {
		// 如果明天才响应, 那么temp += 1 day
		ringDate = temp.Add(time.Hour * 24)
	} else {
		ringDate = temp
	}
	// 休眠时间
	interval := ringDate.Sub(now)

	// 创建定时器channel
	timeChan := getSoleTick(interval)
	// 创建取消channel
	cancelChan := make(chan struct{})
	// 创建跳过一次channel
	jumpOne := make(chan struct{})

	// 开启携程监听channel信息
	go func() {
		nextRing := ringDate
		for {
			select {
			case <-timeChan:
				go d.f()
				timeChan = getSoleTick(time.Hour * 24)
				nextRing.Add(time.Hour * 24)
			case <-cancelChan:
				close(cancelChan)
				close(jumpOne)
				return
			case <-jumpOne:
				nextRing = nextRing.Add(time.Hour * 24)
				interval := nextRing.Sub(time.Now().Local())
				fmt.Println(interval)
				timeChan = getSoleTick(interval)
			case <-cancelChan:
				close(cancelChan)
				return
			}
		}
	}()

	return func() {
			cancelChan <- struct{}{}
		}, func() {
			jumpOne <- struct{}{}
		}
}

func NewDailyScheduler(hour int, min int, second int, callback func()) dailyScheduler {
	return dailyScheduler{hour: hour, min: min, second: second, f: callback}
}
