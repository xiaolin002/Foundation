package test

import (
	"bytes"
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/4 19:25
 **/

// 比较两个切片是否相等

// 采用最原始的挨个遍历的方式
func TestSliceEquals1(t *testing.T) {

	slice1 := make([]int, 0, 10)
	slice2 := make([]int, 0, 11)

	slice1 = append(slice1, 1, 2, 3, 4, 5)
	slice2 = append(slice2, 1, 2, 3, 4, 5)

	// 先比较长度 如果长度不相等 直接返回false
	flag := true
	if len(slice1) != len(slice2) {
		flag = false
	} else {
		for i := 0; i < len(slice1); i++ {
			if slice1[i] != slice2[i] {
				flag = false
				break
			}

		}

	}

	if flag == true && cap(slice1) == cap(slice2) {
		fmt.Println("两个切片相同")
	} else {
		fmt.Println("两个切片不相同")
	}
	return

}

// 采用bytes.equals的方式比较 需要将切片转换成字节切片
func TestSliceEquals2(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}

	fmt.Println(cap(slice1))

	// 将整数切片转换为字节切片
	bytes1 := IntSliceToBytes(slice1)
	bytes2 := IntSliceToBytes(slice2)

	// 使用bytes.Equal比较字节切片
	if bytes.Equal(bytes1, bytes2) {
		fmt.Println("两个切片相同")
	} else {
		fmt.Println("两个切片不相同")
	}
}

// IntSliceToBytes 将整数切片转换为字节切片
func IntSliceToBytes(slice []int) []byte {
	bytes := make([]byte, len(slice)*4) // 假设int类型占用4个字节
	for i, v := range slice {
		bytes[i*4] = byte(v)
		bytes[i*4+1] = byte(v >> 8)
		bytes[i*4+2] = byte(v >> 16)
		bytes[i*4+3] = byte(v >> 24)
	}
	return bytes
}
