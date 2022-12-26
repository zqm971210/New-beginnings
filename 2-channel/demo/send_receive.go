package demo

import "fmt"

func receive(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func SendReceive() {
	ch := make(chan int)
	go receive(ch) // 创建一个 goroutine 从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
