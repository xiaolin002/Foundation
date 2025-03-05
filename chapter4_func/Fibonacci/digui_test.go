package main

import (
	"fmt"
	"math/big"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/14 11:02
 **/

// 递归打印10到1
func TestDiGui10To1(t *testing.T) {
	diGui10To1(10)

}

// 递归打印前30个整数的阶乘的程序
func TestDiGui30(t *testing.T) {
	for i := 0; i <= 30; i++ {
		result := diGui30(i)
		fmt.Printf("%d! = %d\n", i, result)
	}
}

//修正上边的错误 当n=12时会溢出 采用big包修改

func TestDiGui30Big(t *testing.T) {
	for i := 0; i <= 30; i++ {
		result := factorial(int64(i))
		fmt.Printf("%d! = %d\n", i, result)
	}
}

// 计算阶乘的函数
func factorial(n int64) *big.Int {
	result := big.NewInt(1) // 初始化结果为 1
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i)) // result *= i
	}
	return result
}

func diGui30(n int) int {
	if n == 0 {
		return 1
	}
	return n * diGui30(n-1)

}

func diGui10To1(n int) {
	if n > 0 {
		fmt.Println(n)
		diGui10To1(n - 1)
	}

}
