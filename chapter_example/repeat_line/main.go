package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int, 13)

	file, err := os.Open("chapter_example/repeat_line/11.txt")
	if err != nil {
		panic("文件打开失败")
	}
	defer file.Close()

	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d,%s \n", n, line)
		}
	}

}
