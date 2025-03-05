package main

import "fmt"

func main() {
	// 定义一个整数切片
	intSlice := []int{1, 2, 3, 4, 5}

	// 方法一：使用 append 函数复制到空接口切片
	var anySlice1 []interface{}
	for _, v := range intSlice {
		anySlice1 = append(anySlice1, v)
	}
	fmt.Println("使用 append 函数复制的结果:", anySlice1)

	// 方法二：提前分配容量并使用 copy 函数复制到空接口切片
	anySlice2 := make([]interface{}, len(intSlice))
	for i, v := range intSlice {
		anySlice2[i] = v
	}
	fmt.Println("提前分配容量并复制的结果:", anySlice2)
}
