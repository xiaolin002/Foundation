package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// 统计了 Unicode 字符的出现次数和 UTF-8 编码长度的分布情况，并报告了无效 UTF-8 字符的数量
func main() {

}

// 字符出现的次数
func charCount() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}

// 各种字符出现次数
func charAllCount() {
	letterCounts := make(map[rune]int) // counts of 字母
	digitCounts := make(map[rune]int)  // counts of 数字
	spaceCounts := make(map[rune]int)  // counts of 空格字符 表示的'\t', '\n', '\v', '\f', '\r', ' ' 这些类型
	otherCounts := make(map[rune]int)  // counts of 其他类型的字符
	var utflen [utf8.UTFMax + 1]int    // count of lengths of UTF-8 encodings
	invalid := 0                       // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCounts[r]++
		} else if unicode.IsDigit(r) {
			digitCounts[r]++
		} else if unicode.IsSpace(r) {
			spaceCounts[r]++
		} else {
			otherCounts[r]++
		}
		utflen[n]++
	}
	fmt.Printf("letter\tcount\n")
	for c, n := range letterCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\ndigit\tcount\n")
	for c, n := range digitCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nspace\tcount\n")
	for c, n := range spaceCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nother\tcount\n")
	for c, n := range otherCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}

func wordCount() {
	// 创建一个 map 来存储每个单词的出现次数
	wordFreq := make(map[string]int)

	// 从标准输入读取数据
	scanner := bufio.NewScanner(os.Stdin)

	// 设置分割函数为按单词分割
	scanner.Split(bufio.ScanWords)

	// 遍历输入的单词
	for scanner.Scan() {
		word := scanner.Text()
		// 统计单词出现的次数
		wordFreq[word]++
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	// 输出每个单词及其出现的频率
	fmt.Println("Word\tFrequency")
	for word, freq := range wordFreq {
		fmt.Printf("%s\t%d\n", word, freq)
	}

}
