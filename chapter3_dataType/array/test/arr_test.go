package test

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/1 22:09
 **/

func TestArr(t *testing.T) {

	//定义了一个含有100个元素的数组r，最后一个元素被初始化为-1，其它元素都是用0初始化
	r := [...]int{99: -1}
	for _, v := range r {
		fmt.Printf("%d\n", v)
	}
	fmt.Println(len(r))
}

func TestArrEquals(t *testing.T) {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"

	// 这里IDE会报错 因为数组长度不同
	// d := [3]int{1, 2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}

func TestSha256(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func TestSha256Equals(t *testing.T) {
	count := CountDifferentBits(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X")))
	fmt.Println(count)
}

func CountDifferentBits(c1, c2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			// 这里的操作涉及到sha256的原理  这里的1代表的是一个字节  8个字节代表一个数字
			// 了解到两个数组如何对比即可
			bit1 := (c1[i] >> j) & 1
			bit2 := (c2[i] >> j) & 1
			if bit1 != bit2 {
				count++
			}
		}
	}
	return count
}
func zeroPtr(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// 换一种写法

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}

}
