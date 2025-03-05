package main

import (
	"fmt"
)

func main() {

	var sum = "123456"
	var s, sep string
	var result []string
	for i := 0; i < len(sum); i++ {
		s = s + sep + string(sum[i])
		sep = " "
		result = append(result, s)
	}
	fmt.Println(s)
	// 打印拼接后的结果
	for _, str := range result {
		fmt.Println(str)
	}

}
