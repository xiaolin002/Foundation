package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改值
// 通过CanSet 判断是否可以修改
// 通过Elem 方法获取到指针指向的值
func main() {

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//panic: reflect: reflect.Value.SetFloat using unaddressable value
	//v.SetFloat(7.1)
	fmt.Println("settability of v:", v.CanSet())

	// 必须传递指针
	fmt.Println("--------传递指针--------------:")
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet()) // 仍然是false
	fmt.Println("---------使用elem方法------------:")
	v = v.Elem()
	fmt.Println("settability of v:", v.CanSet()) // 现在是true
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(v)

}
