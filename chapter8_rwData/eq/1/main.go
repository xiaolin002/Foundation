package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var charCount, wordCount, lineCount int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("请输入文本（输入 'S' 结束）：")

	for {
		// 读取一行输入
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入时出错:", err)
			return
		}

		// 去掉行末的换行符
		input = strings.TrimSuffix(input, "\n")

		// 检查是否输入 'S'，如果是则结束输入
		if input == "S" {
			break
		}

		// 统计字符数（包括空格）
		charCount += len(input)

		// 统计单词数
		words := strings.Fields(input)
		wordCount += len(words)

		// 统计行数
		lineCount++
	}

	// 输出结果
	fmt.Printf("字符个数（包括空格）：%d\n", charCount)
	fmt.Printf("单词个数：%d\n", wordCount)
	fmt.Printf("行数：%d\n", lineCount)
}
