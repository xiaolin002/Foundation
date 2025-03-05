package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	newElements := []int{10, 20, 30}

	// 先删除要替换的元素 后添加一个空切片
	slice = append(slice[:2], slice[5:]...)
	fmt.Println(slice) // 输出: [1 2 ]

	// 再添加新的元素
	slice = append(slice, newElements...)

	fmt.Println(slice) // 输出: [1 2 10 20 30]

}
