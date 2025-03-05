package main

import "fmt"

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

// 实现 Shaper 接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shaper

	// 创建 Circle 和 Rectangle 实例
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 3, Height: 4}

	// 测试 Circle 是否实现了 Shaper 接口
	if _, ok := interface{}(circle).(Shaper); ok {
		fmt.Println("Circle 实现了 Shaper 接口")
	} else {
		fmt.Println("Circle 未实现 Shaper 接口")
	}

	/*
		这样写也可以
			var s Shaper = circle
			if _, ok := s.(Shaper); ok {
			    fmt.Println("Circle 实现了 Shaper 接口")
			} else {
			    fmt.Println("Circle 未实现 Shaper 接口")
			}

	*/

	// 测试 Rectangle 是否实现了 Shaper 接口
	if _, ok := interface{}(rectangle).(Shaper); ok {
		fmt.Println("Rectangle 实现了 Shaper 接口")
	} else {
		fmt.Println("Rectangle 未实现 Shaper 接口")
	}

	/*
		这样封装起来也行
		func implementsShaper(v interface{}) bool {
		    _, ok := v.(Shaper)
		    return ok
		}

		if implementsShaper(circle) {
		    fmt.Println("Circle 实现了 Shaper 接口")
		} else {
		    fmt.Println("Circle 未实现 Shaper 接口")
		}

	*/

	// 赋值给接口变量
	s = circle
	fmt.Printf("圆的面积: %.2f\n", s.Area())

	s = rectangle
	fmt.Printf("矩形的面积: %.2f\n", s.Area())
}
