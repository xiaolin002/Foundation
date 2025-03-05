// E:/go_example/foundation/chapter6_interface/simple/main.go
package main

import "fmt"

// Simpler 定义接口
type Simpler interface {
	Get() int
	Set(int)
}

// Simple 定义结构体类型
type Simple struct {
	value int
}

// Get 实现 Simpler 接口的 Get 方法
func (s *Simple) Get() int {
	return s.value
}

// Set 实现 Simpler 接口的 Set 方法
func (s *Simple) Set(v int) {
	s.value = v
}

// operate 定义一个函数，接收 Simpler 类型的参数
func operate(s Simpler) {
	// 调用 Set 方法设置值
	s.Set(42)
	// 调用 Get 方法获取值
	result := s.Get()
	fmt.Println("The value is:", result)
}

func main() {
	// 创建 Simple 结构体的实例
	s := &Simple{}
	// 调用 operate 函数
	operate(s)
}
