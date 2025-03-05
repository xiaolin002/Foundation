package main

import "fmt"

// myPrintInterface 定义一个接口，包含 print 方法
type myPrintInterface interface {
	print()
}

// extendedPrintInterface 定义一个新接口，嵌套 myPrintInterface
type extendedPrintInterface interface {
	myPrintInterface
	extendedPrint()
}

// MyStruct 定义一个结构体，实现 extendedPrintInterface 接口
type MyStruct struct{}

// 实现 myPrintInterface 接口的 print 方法
func (m MyStruct) print() {
	fmt.Println("执行 print 方法")
}

// 实现 extendedPrintInterface 接口的 extendedPrint 方法
func (m MyStruct) extendedPrint() {
	fmt.Println("执行 extendedPrint 方法")
}

// f3 函数接受一个 myPrintInterface 类型的参数
func f3(x myPrintInterface) {
	x.print()
}

func main() {
	// 创建 MyStruct 实例
	var s MyStruct

	// 声明一个 extendedPrintInterface 类型的变量，并赋值为 s
	var epi extendedPrintInterface = s

	// 因为 extendedPrintInterface 嵌套了 myPrintInterface，所以可以将 epi 赋值给 myPrintInterface 类型的变量
	var mpi myPrintInterface = epi

	// 调用 f3 函数，传入 mpi
	f3(mpi)

	// 调用 extendedPrint 方法
	epi.extendedPrint()
}
