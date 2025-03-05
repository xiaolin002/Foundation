package main

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/15 11:01
 **/

// 测试当数组赋值时 是否发生了数组内存拷贝
func TestCopyArr(t *testing.T) {

	var arr = [5]int{1, 2, 3, 4, 5}
	var arr2 = arr
	arr2[0] = 100
	fmt.Println(arr, arr)
	fmt.Println(arr2, arr2)
	//[1 2 3 4 5] &[1 2 3 4 5]
	//[100 2 3 4 5] &[100 2 3 4 5]
}

// 采用数组的方式计算出斐波那契数列的前50个数
func TestFibonacci(t *testing.T) {
	var arr [50]int
	arr[0] = 1
	arr[1] = 1
	//fib := [50]int{0, 1}  这种方式也行
	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	fmt.Println(arr)
}

// 当数组用【...】表示时，其本质上已经变成了切片
func TestLenArr(t *testing.T) {
	var arrLazy = [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(arrLazy), cap(arrLazy))
}
