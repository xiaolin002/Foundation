package main

import (
	"fmt"
	"math/rand"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/11 16:15
 **/

func TestRandStr(t *testing.T) {
	strings := make([]string, 16)

	// 定义一个包含数字和字母的字符集
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 遍历字符串切片，为每个元素生成随机字符串
	for i := range strings {
		// 创建一个长度为8的随机字符串
		randomString := make([]byte, 8)
		for j := 0; j < 8; j++ {
			randomString[j] = charset[rand.Intn(len(charset))]
		}
		// 将随机字符串添加到切片中
		strings[i] = string(randomString)
	}

	// 打印生成的随机字符串切片
	fmt.Println(strings)

}
