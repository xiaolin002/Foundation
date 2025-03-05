package main

import (
	"fmt"
	"time"
)

// Future 结构体表示一个未来结果
type Future struct {
	resultChan chan int
}

// Get 方法用于获取未来结果
func (f *Future) Get() int {
	return <-f.resultChan
}

// compute 函数模拟一个耗时的计算任务
func compute() *Future {
	future := &Future{
		resultChan: make(chan int),
	}

	go func() {
		// 模拟耗时操作
		time.Sleep(2 * time.Second)
		result := 42
		future.resultChan <- result
	}()

	return future
}

func main() {
	// 启动计算任务并获取 Future 对象
	future := compute()

	// 继续执行其他任务
	fmt.Println("Doing other work...")

	// 获取计算结果
	result := future.Get()
	fmt.Printf("The result is: %d\n", result)
}
