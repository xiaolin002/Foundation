package main

import (
	"fmt"
	"time"
)

func main() {
	//当设置成有缓冲的通道时，1秒可以接收很多数据。
	ch1 := make(chan int, 100)
	go pump(ch1)
	go push(ch1)
	time.Sleep(1e9)

}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func push(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
