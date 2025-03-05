package main

import (
	"fmt"
)

// 通道工厂函数，返回一个用于计算两数之和的通道
func sumChannelFactory(a, b int) chan int {
	// 创建一个无缓冲的整数通道
	resultChan := make(chan int)
	// 启动一个协程来计算两数之和
	go func() {
		// 计算两数之和
		sum := a + b
		// 将计算结果发送到通道中
		resultChan <- sum
		// 关闭通道
		close(resultChan)
	}()
	return resultChan
}

func main() {
	num1 := 10
	num2 := 20
	// 调用通道工厂函数获取用于计算两数之和的通道
	resultChan := sumChannelFactory(num1, num2)
	// 等待协程计算结果并从通道中接收
	result := <-resultChan
	// 打印计算结果
	fmt.Printf("%d + %d = %d\n", num1, num2, result)
}
