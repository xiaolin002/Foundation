package main

import (
	"fmt"
	"time"
)

func doWork() chan string {
	result := make(chan string)
	go func() {
		// 模拟一个耗时操作
		time.Sleep(2 * time.Second)
		result <- "Work done"
	}()
	return result
}

func main() {
	resultChan := doWork()
	select {
	case result := <-resultChan:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}
}
