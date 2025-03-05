package main

import (
	"fmt"
	"testing"
)

/**
 * @Description  接口断言
 * @Date 2025/2/15 11:48
 **/

type Shapper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Circler struct {
	radius float32
}

func (ci *Circler) Area() float32 {
	return ci.radius * ci.radius * 3.14
}

func TestShaper(t *testing.T) {
	// 类型断言
	var shaper Shapper
	sq1 := new(Square)
	sq1.side = 5
	shaper = sq1

	// 注意这里一定用的是指针类型   如过Area是值类型接收者  这里就不能用指针类型
	// 否则会报错
	if value, ok := shaper.(*Square); ok {
		fmt.Println(value)
		fmt.Println("Square Area is:", value.Area())
	}

	if value, ok := shaper.(*Circler); ok {
		fmt.Println("Circle Area is:", value.Area())
	} else {
		fmt.Println("interface not has a struct ")
	}

}
