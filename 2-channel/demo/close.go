package demo

import (
	"fmt"
	"time"
)

func Close() {
	go func() {
		time.Sleep(time.Hour)
	}()
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	c <- 3
}

func Close2() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	fmt.Println(<-c) //1
	fmt.Println(<-c) //2
	fmt.Println(<-c) //0
	fmt.Println(<-c) //0
}

func Close3() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}

func Close4() {
	c := make(chan int, 10)
	close(c)
	i, ok := <-c
	fmt.Printf("%d, %t", i, ok) //0, false
}
