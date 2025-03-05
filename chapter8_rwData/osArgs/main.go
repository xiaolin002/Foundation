package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 这里用控制台 go build main.go
	// 然后go run main.go 后面加参数  如果参数有多个 用空格隔开
	// 比如 go run main.go 1 2 3 4 5
	// os.Args[0]放的是程序本身的名字

	who := "Alice--"
	if len(os.Args) > 1 {
		// strings.Join 将多个字符串连接起来，中间是空格
		who += strings.Join(os.Args[1:], "--")
	}
	fmt.Println("Good morning", who)
	fmt.Println("-----》", os.Args[0])

}
