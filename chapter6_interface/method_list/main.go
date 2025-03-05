package main

import "fmt"

type Appender interface {
	Append(int)
}
type Lener interface {
	Len() int
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	//声明一个指向 List 类型的指针
	var lst List
	// 把 lst 作为 Appender 和 Lener 类型的参数使用
	// lst 非指针类型	报错
	//CountInto(lst, 1, 10)

	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	// 把一个指向 lst 的指针作为 Appender 类型的参数使用
	plst := new(List)
	CountInto(plst, 1, 10)
	// 指针被解引用
	if LongEnough(plst) {
		// lst 现在的长度是 10
		fmt.Printf("- plst is long enough\n")
	}
	//var apender Appender = plst // 正确，*List 实现了 Appender 接口
	//var appender Appender = lst // 错误，List 未实现 Appender 接口
}

/*

接受者是值的方法可以通过值和指针调用，因为指针会首先被解引用
是接受者是指针的方法只能通过指针调用，因为存储在接口中的值没有地址

值类型：若类型 T 有以 T 为接收者的方法，这些方法属于 T 的方法集；若类型 T 有以 *T 为接收者的方法，这些方法不属于 T 的方法集，但可以通过 T 的实例调用（Go 会自动取地址）。
指针类型：若类型 *T 有以 T 或 *T 为接收者的方法，这些方法都属于 *T 的方法集

这意味着只有 *List 类型（List 类型的指针）才实现了 Appender 接口，而 List 类型本身并未实现。所以，当你尝试把 lst（List 类型）传入 CountInto 函数时，编译器会报错，因为 lst 不是 Appender 接口的实现类型。
lst 能执行 Append 方法的原因
在 Go 语言中，当调用使用指针接收者定义的方法时，若接收者是值类型，Go 会自动取其地址。所以lst.Append(1)，尽管 Append 方法是用指针接收者定义的，但你依然可以通过 lst.Append(1) 调用该方法，Go 会自动将其转换为 (&lst).Append(1)
*/
