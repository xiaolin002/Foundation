package main

import (
	"fmt"
	"strings"
)

func main() {
	//
	//var str string = "goWeb is good"
	//str = "Hi" + str[5:]
	//fmt.Println(str)
	//
	//var str1 string = "goWeb is good"
	//str2 := []byte(str1)
	//str2[0] = 'G'
	//str1 = string(str2)

	original := "Hello, world! Hello, Go!"
	// n表示替换个数
	replaced := strings.Replace(original, "Hello", "Hi", 3)
	fmt.Println(replaced) // 输出：Hi, world! Hi, Go!

}
