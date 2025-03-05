package main

import (
	"fmt"
	"reflect"
)

// 注意结构体中的字段必须是可导出(大写)的，才能被反射访问被修改

type T struct {
	Age     int
	Address string
}

func main() {
	t := T{23, "skidoo"}
	// 注意一定传递指针 然后使用Elem方法 最后用CanSet判断是否可以修改
	s := reflect.ValueOf(&t).Elem()
	fmt.Println(s.CanSet())

	t2 := s.Type()
	for i := 0; i < s.NumField(); i++ {
		// 获取每个字段的值
		f := s.Field(i)
		//输出具体的字段名称和值
		fmt.Printf("%d: %s %s = %v\n", i, t2.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)

	///不写 t2 := s.Type() 直接获取字段名称是不行的。
	//因为 s 是 reflect.Value 类型，该类型主要用于操作值，
	//虽然可以通过 s.Field(i) 获取字段的值，但无法直接获取字段的名称。
	//而 reflect.Type 类型则提供了获取类型信息的方法，
	//像 Field(i).Name 就可以用来获取结构体中第 i 个字段的名称

}
