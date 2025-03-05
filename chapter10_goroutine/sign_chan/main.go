package main

import (
	"fmt"
	"time"
)

// worker 函数模拟一个处理任务，完成后向通道发送信号
func worker(ch chan bool) {
	fmt.Println("第一个协程开始处理工作...")
	time.Sleep(3 * time.Second)
	fmt.Println("第一个协程处理工作完成。")
	ch <- true
}

// 新增的 worker 类型函数
func anotherWorker(ch chan bool) {
	fmt.Println("第二个协程开始处理工作...")
	time.Sleep(2 * time.Second)
	fmt.Println("第二个协程处理工作完成。")
	ch <- true
}

func main() {
	// 创建一个布尔类型的通道
	ch := make(chan bool)

	// 启动两个协程
	go worker(ch)
	go anotherWorker(ch)

	// 主协程等待两个协程处理结束
	fmt.Println("主协程等待协程处理结束...")
	// 等待第一个协程完成
	<-ch
	// 等待第二个协程完成
	<-ch
	fmt.Println("主协程接收到所有协程处理结束的信号，程序结束。")
}
