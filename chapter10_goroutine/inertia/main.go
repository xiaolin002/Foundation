package main

import (
	"fmt"
)

// naturalNumbers 是一个惰性生成器函数，用于生成自然数序列
func naturalNumbers() chan int {
	// 创建一个无缓冲的整数通道
	numbers := make(chan int)
	go func() {
		// 从 0 开始生成自然数
		for i := 0; ; i++ {
			// 将生成的自然数发送到通道中
			numbers <- i
		}
	}()
	return numbers
}

func main() {
	// 调用 naturalNumbers 函数获取生成器通道
	numChan := naturalNumbers()
	// 只打印前 10 个自然数
	for i := 0; i < 10; i++ {
		// 从通道中接收并打印自然数
		fmt.Println(<-numChan)
	}
}
