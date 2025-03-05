package main

import (
	"fmt"
	"reflect"
)

// 例子用来学习
func main() {

	var x float64 = 3.4
	fmt.Println("type :", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	//判断是float 32 还是64
	fmt.Println("value", v.Float())
	// 输出 reflect.Value 转换回的 interface{} 类型的值
	fmt.Println(v.Interface())
	// 尝试将 interface{} 类型的值转换为 float64 类型
	y, ok := v.Interface().(float64)
	if !ok {
		fmt.Println("类型断言失败")
	}

	fmt.Println(y)

}
