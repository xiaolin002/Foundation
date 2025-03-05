package main

import (
	"fmt"
	"reflect"
)

// 反射结构体
// NumField 返回结构体的字段数量
// Field 返回结构体的字段

type NotKnownType struct {
	s1, s2, s3 string
}

func (n NotKnownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3

}

var secret interface{} = NotKnownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(value)
	fmt.Println(typ) //main.NotKnownType
	knd := value.Kind()
	fmt.Println(knd) // struct

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("field %d: %v\n", i, value.Field(i))
	}

	//results := value.Method(0).Call(nil)
	//fmt.Println(results)

	//优化
	// 检查第 0 个方法是否存在
	if value.NumMethod() > 0 {
		method := value.Method(0)
		// 调用方法并传递空参数列表
		results := method.Call(nil)
		// 检查返回值是否存在
		if len(results) > 0 {
			// 提取第一个返回值并进行类型转换
			resultStr := results[0].String()
			fmt.Println(resultStr)
		}
	} else {
		fmt.Println("该结构体没有定义任何方法")
	}

}
