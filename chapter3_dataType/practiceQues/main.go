package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"
	fmt.Println("Length in bytes:", len(str)) // 汉字占3个字节
	fmt.Println("Length in runes:", utf8.RuneCountInString(str))

}
