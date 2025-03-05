package main

import "fmt"

type Sorter interface {
	// Len 长度
	Len() int
	// Less 比较大小
	Less(i, j int) bool
	// Swap 交换位置
	Swap(i, j int)
}

type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort 冒泡排序
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}
func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data)
	Sort(a)
	fmt.Printf("%v\n", a)

}
