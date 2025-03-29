package main

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/3/29 10:09
 **/

func TestDeferV0(t *testing.T) {
	fmt.Println(V0())
}

func V0() int {
	i := 0
	defer func() {
		i = 1
	}()
	// 输出0
	return i
}

func TestDeferV1(t *testing.T) {
	V1()

}
func V1() {
	i := 0
	defer func(val int) {
		println(i)
	}(i)
	// 输出1
	i = 1
}

func TestDeferV2(t *testing.T) {
	V2()

}

func V2() {
	i := 0
	defer func() {
		// 输出1
		println(i)
	}()
	i = 1
}

func TestDeferV3(t *testing.T) {
	v3 := V3()
	fmt.Println(v3)
}

func V3() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	// 函数有明确的返回值，defer 语句会在函数返回前执行，
	// 因此 defer 语句中的赋值操作会覆盖函数的返回值。
	return a
}

func TestDeferV4(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

}

func TestDeferV5(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Println(val)
		}(i)
	}
}

func TestDeferV6(t *testing.T) {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Println(j)
		}()
	}

}
