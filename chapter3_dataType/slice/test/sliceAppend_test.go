package test

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/4 21:13
 **/

func TestAppend(t *testing.T) {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// len大于 cap 需要扩容
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 扩容后就不在使用原数组了
	}
	z[len(x)] = y // 将y添加到起切片中
	return z      // 返回一个新的切片
}

func appendMoreInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, eleven_1 amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z[len(x):], y)

	}
	return z
}

func TestNonempty(t *testing.T) {
	arr := []string{"a", "b", "c", "", "d"}
	result := nonempty(arr)
	fmt.Println(result)

}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {

		// 当s为空的字符窜的时候if条件不执行 i不会变然后执行下一个s
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
