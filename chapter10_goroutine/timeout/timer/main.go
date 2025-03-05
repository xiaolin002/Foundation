package main

import (
	"fmt"
	"time"
)

func doWorkWithTimer(timer *time.Timer) string {
	resultChan := make(chan string)
	go func() {
		// 模拟一个耗时操作
		time.Sleep(2 * time.Second)
		resultChan <- "Work done"
	}()

	select {
	case result := <-resultChan:
		timer.Stop()
		return result
	case <-timer.C:
		return "Timeout!"
	}
}

func main() {
	timer := time.NewTimer(1 * time.Second)
	result := doWorkWithTimer(timer)
	fmt.Println(result)
}
