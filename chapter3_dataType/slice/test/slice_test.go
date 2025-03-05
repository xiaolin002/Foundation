package test

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/4 17:07
 **/

//结论是 如果切片是[5:] 则长度和容量就是原来切片的长度和容量减去5
//      如果切片是[:5] 则长度和容量都还是原来切片的长度和容量
//  结论 从头去取则长度容量不变  从中间取则长度和容量都变了 未发生扩容还是用的底层数组

func TestSliceMonth(t *testing.T) {
	months := [...]string{1: "January", 2: "February", 3: "March",
		4: "April", 5: "May", 6: "June", 7: "July",
		8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(cap(months)) //13

	bbb := months[5:]
	ccc := months[:5]
	fmt.Printf("bbb值是%s,cap容量%d,len长度是%d\n", bbb, cap(bbb), len(bbb))
	fmt.Printf("ccc值是%s,cap容量%d,len长度是%d\n", ccc, cap(ccc), len(ccc))

	fmt.Println("----------------------------")

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				//June
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	aaa := Q2[:5]
	fmt.Println(aaa)      //[April May June July August]
	fmt.Println(cap(aaa)) // 9
	// 这里本质上summer的底层数组和Q2的底层数组是同一个 都是months
	endlessSummer := summer[:5]     // extend a slice (within capacity)
	fmt.Println(endlessSummer)      //[June July August September October]
	fmt.Println(cap(endlessSummer)) // 7

}

func TestReversSlice(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	reverseSlice(a[:])
	fmt.Println(a)
}

func reverseSlice(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

}

// 重写reverse函数，使用数组指针代替slice
func TestReversSlice2(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	reverseSlice2(&a)
	fmt.Println(a)
}

// 注意如果使用数组指针作为形参的话 则形参的长度是固定的 与传递数组的长度一致
func reverseSlice2(arr *[7]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func TestReversNLeft(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverseSlice(s[:2])
	reverseSlice(s[2:])
	reverseSlice(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	// 理解为向左推2个数

}
func TestReversNRight(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverseSlice(s)
	reverseSlice(s[:2])
	reverseSlice(s[2:])

	fmt.Println(s) //
}

// 去除相邻之间重复的字符串	aabbccdd => abcd
func dedup(strings *[]string) {
	if len(*strings) == 0 {
		return
	}

	// Use two pointers to traverse the slice
	i, j := 0, 1
	for j < len(*strings) {
		if (*strings)[i] != (*strings)[j] {
			i++
			(*strings)[i] = (*strings)[j]
		}
		j++
	}

	// Truncate the slice to remove duplicates
	*strings = (*strings)[:i+1]
}

func TestDedup(t *testing.T) {
	strings := []string{"a", "b", "b", "c", "c", "c", "d", "d", "d", "d"}
	dedup(&strings)
	fmt.Println(strings) // Output: [a b c d]
}

// 声明一个切片
func TestSlice(t *testing.T) {
	s := [3]int{1, 2, 3}
	fmt.Println(s[:])

}

func TestByte(t *testing.T) {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(string(b[1:4]))
	fmt.Println(string(b[:2]))
	fmt.Println(string(b[:]))
}

func TestSliceSplite(t *testing.T) {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Println(slice1)
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

}

// 切片传递
func TestSliceTra(t *testing.T) {

}

func Sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s

}

func appendByte(slice []byte, data []byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, eleven_1 future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
