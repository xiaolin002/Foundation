package main

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/14 18:20
 **/

func TestClosure(t *testing.T) {

	func() {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
			t.Logf("sum = %d", sum)
		}
	}()

}

func TestClosure1(t *testing.T) {
	g := func(i int) {
		t.Logf("i = %d", i)
	}
	func() {
		for i := 0; i < 4; i++ {
			g(i)
			fmt.Printf("-g is of type %T ", g) // 添加换行符
		}

	}()

}

func TestHelloWord(t *testing.T) {
	a := func() {
		fmt.Println("hello world")

	}
	a()

}

func TestReturnDefer(t *testing.T) {
	fmt.Println(f())
	// 输出2  先执行return 再执行defer

}

func f() (res int) {
	defer func() {
		res++
	}()
	return 1

}

func TestAddr(t *testing.T) {
	var f1 = Adder()
	fmt.Print(f1(1), " - ")
	fmt.Print(f1(20), " - ")
	fmt.Print(f1(300), " - ")

}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}

}
