package main

import "fmt"

func splitString(str string, i int) (string, string) {
	if i < 0 || i >= len(str) {
		return "", ""
	}
	return str[:i], str[i:]
}

func main() {
	str := "Hello, 世界"
	i := 7
	part1, part2 := splitString(str, i)
	fmt.Printf("原始字符串: %s\n", str)
	fmt.Printf("分割索引: %d\n", i)
	fmt.Printf("分割后的字符串1: %s\n", part1)
	fmt.Printf("分割后的字符串2: %s\n", part2)

	str2 := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(str2[len(str2)/2:] + str2[:len(str2)/2])
}
