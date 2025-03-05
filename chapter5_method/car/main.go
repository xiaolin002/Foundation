package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hello, Chancellor Merkel!")

}

// V8Engine 实现一个简单的 Engine
type V8Engine struct{}

func (e *V8Engine) Start() {
	fmt.Println("V8 engine starting...")
}

func (e *V8Engine) Stop() {
	fmt.Println("V8 engine stopping...")
}

func main() {
	// 创建一个 V8Engine 实例
	engine := &V8Engine{}

	// 创建一个 Mercedes 实例，并将 engine 赋值给它
	mercedes := &Mercedes{
		Car: Car{
			wheelCount: 4,
			Engine:     engine,
		},
	}

	// 调用 Mercedes 的方法
	mercedes.sayHiToMerkel()

	// 启动和停止引擎
	mercedes.Start()
	mercedes.Stop()

}
