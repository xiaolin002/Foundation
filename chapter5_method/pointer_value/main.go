package main

import "fmt"

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}

func (b B) write() string {
	b.thing = 2
	return fmt.Sprint(b)
}

func main() {
	var b1 B                // b1是值
	b1.change()             // 这里change是指针类型 默认将b的地址传入
	fmt.Println(b1.write()) //【2】这里是值类型 所以不会改变b1的值 相当于在这里创建了一个副本 b3 := b1 然后在b3上操作最后输出b3的值是2
	fmt.Println(b1.thing)   // 输出1

	b2 := new(B)
	b2.change()             // 这里change是指针类型 默认将b的地址传入
	fmt.Println(b2.write()) // 输出{2}
	fmt.Println(b2.thing)   // 输出1

}
