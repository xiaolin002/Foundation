package main

import (
	"fmt"
	"strings"
)

func main() {
	addBmp := makeAddSuffix(".bmp")
	addJpeg := makeAddSuffix(".jpeg")

	fmt.Println(addBmp("file"))
	fmt.Println(addJpeg("file"))

}

// 在返回的函数中，使用 strings.HasSuffix 检查 name 是否已经以 suffix 结尾。
// 如果没有，则将 suffix 添加到 name 的末尾并返回；如果已经有后缀，则直接返回 name。
func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix

		}
		return name
	}

}
