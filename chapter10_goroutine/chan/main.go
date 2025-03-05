package main

import (
	"fmt"
)

func main() {

	var ch chan string
	// 记得管道要初始化
	ch = make(chan string)

	go sendData(ch)

	go getData(ch)

}

func sendData(ch chan string) {
	ch <- "W3Cschool教程"
	ch <- "W3Cschool"
	ch <- "W"
	ch <- "你好"
	ch <- "你不好"
	ch <- "你真"
	ch <- "你帮"
}

func getData(ch chan string) {
	for {
		input := <-ch
		fmt.Printf("%s\n", input)

	}
}
