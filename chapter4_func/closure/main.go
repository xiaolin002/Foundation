package main

import "fmt"

// 使用闭包的方式实现斐波那契数列
// 闭包的实现方式是将函数作为返回值，同时将函数内部的变量作为返回值
func main() {

	f := fibonacci()
	for i := 0; i < 10; i++ {
		//这里调用10次 打印出前10个斐波那契数列
		fmt.Println(f())
	}

}

func fibonacci() func() int {

	// 这里就相当于定义当 n=1时，值为1
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

// 动态规划计算斐波那契数列
func fibonacci_dp(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 迭代计算斐波那契数列
func fibonacci_iterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
