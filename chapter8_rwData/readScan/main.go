package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName string
)

func main() {
	//fmt.Println("please input your full name")
	//// 用空格分隔输入
	//fmt.Scanln(&firstName, &lastName)
	//fmt.Println("your full name is", firstName, lastName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your full name")
	input, err := reader.ReadString('\n') // delim 是碰到指定的字符就停止
	if err != nil {
		fmt.Println("read input error", err)
		return
	}
	fmt.Print("your full name is", input)

}
