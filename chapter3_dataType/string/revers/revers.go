package revers

/**
 * @Description
 * @Date 2025/1/16 17:47
 **/

// 两种写法  一种是直接用rune  一种是直接用byte
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func reverseStringInPlace(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 反转字符串，使用两个切片
func reverseStringWithTwoSlices(s string) string {
	runes := []rune(s)
	reversed := make([]rune, len(runes))
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		reversed[i] = runes[j]
	}
	return string(reversed)
}

// 反转字符串，使用一个切片
func reverseStringWithOneSlice(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 写一个程序 要求能够遍历一个数组的字符 并将当前字符和前一个字符不相同的字符拷贝到另一个数组当中
func copyUniqueChars(input []byte) []byte {
	if len(input) == 0 {
		return []byte{}
	}

	// 创建一个新的切片来存储结果
	result := make([]byte, 0, len(input))

	// 将第一个字符添加到结果切片中
	result = append(result, input[0])

	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] {
			result = append(result, input[i])
		}
	}

	return result
}
