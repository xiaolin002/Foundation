package main

import "fmt"

/*
参照sorter接口 创建一个Miner接口实现一些必要的操作
函数min接受一个Miner类型变量的集合
然后计算病返回集合中最小的元素
*/

type Miner interface {
	Len() int
	Less(i, j int) bool
	Get(i int) any // 获取指定索引的元素
}

type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Get(i int) any      { return p[i] }

func Min(data Miner) any {
	if data.Len() == 0 {
		fmt.Println("data is empty")
		return nil
	}
	mi := 0

	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			mi = i
		}
	}
	return data.Get(mi)

}

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data)
	// 调用 min 函数找出最小元素
	result := Min(a)
	fmt.Printf("集合中最小的元素是: %v\n", result)

}
