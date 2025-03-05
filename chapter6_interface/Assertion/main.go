package main

import "fmt"

/*
   类型推断
*/

// Shaper 定义一个接口
type Shaper interface {
	Area() float64
}

// Circle 定义一个结构体
type Circle struct {
	Radius float64
}

// 实现 Shaper 接口
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle 定义另一个结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现 Shape 接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shaper) {
	// 这里的参数如果改为是空接口类型 s interface{}，就需要断言
	//if shaper, ok := s.(Shaper); ok {
	//		switch v := shaper.(type) {
	switch v := s.(type) {
	case Circle:
		fmt.Printf("圆的面积是: %.2f\n", v.Area())
	case Rectangle:
		fmt.Printf("矩形的面积是: %.2f\n", v.Area())
	default:
		fmt.Println("未知形状")
	}
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 3, Height: 4}

	printArea(circle)
	printArea(rectangle)
}
