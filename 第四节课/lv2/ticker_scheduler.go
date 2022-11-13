package main

import (
	"time"
)

type tickerScheduler struct {
	t time.Duration
	f func()
}

func (ts tickerScheduler) Run() func() {
	ticker := time.NewTicker(ts.t)
	cancel := make(chan struct{})

	go func() {
		// 捕获了ts, tiker和cancel
		for {
			select {
			case <-ticker.C:
				go ts.f()
			case <-cancel:
				ticker.Stop()
				close(cancel)
				return
			}
		}
	}()

	return func() {
		cancel <- struct{}{}
	}
}

func NewTickerScheduler(interval time.Duration, callback func()) tickerScheduler {
	return tickerScheduler{t: interval, f: callback}
}
