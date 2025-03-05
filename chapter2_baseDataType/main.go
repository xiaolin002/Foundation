package main

import (
	"fmt"
	ft "foundation/chapter2_baseDataType/type"
)

type Celsius float64

func main() {
	var c Celsius = 37.0
	f := Fahrenheit(c) // 类型转换
	// 隐式调用了其中的String方法
	fmt.Println(f)

	fmt.Println("-----------------------------------")
	// 创建 FT 类型的实例
	ftInstance := ft.FT(ft.MyFunction)

	// 调用 ftInstance 函数
	_, err := ftInstance("a", 2)
	if err != nil {
		fmt.Println("An error occurred:", err)
	} else {
		fmt.Println("Function call was successful")
	}

	fmt.Println("----------------------------------")

	f2 := new(ft.Stack)
	f2.Push(1)
	// 也可以写成
	f3 := &ft.Stack{}
	f3.Push(3)

}

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}
