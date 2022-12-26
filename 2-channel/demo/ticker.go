package demo

import (
	"fmt"
	"time"
)

// ticker是一个定时触发的计时器
func Ticker() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Second * 3)
}
