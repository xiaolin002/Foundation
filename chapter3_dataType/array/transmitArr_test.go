package main

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/15 11:36
 **/

// 数组传递时必须指定长度（不论传递时还是返回时）
// 这里的sum不指定长度的话就是切片切片就得初始化 否则报错
// 否则就在下边通过make函数初始化
//大多数的话还是用切片传递

func Sum(a *[3]float64) (sum [3]float64) {
	a[0] = 1
	a[1] = 2
	a[2] = 3

	//sum = make([]float64, 3)
	// 这个是数组的传递方法
	for i := 0; i < len(a); i++ {
		sum[i] = a[i]

	}

	// 这个是切片的传递方法
	//sum = append(sum, a[:]...)
	return
}

func TestArr(t *testing.T) {
	array := [3]float64{1.1, 2.2, 3.3}
	x := Sum(&array)
	fmt.Println(x)
}
