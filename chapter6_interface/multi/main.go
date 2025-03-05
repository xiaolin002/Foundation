package main

import "fmt"

// 接口的提取

type Shaper interface {
	Area() float32
}
type TopologicalGenus interface {
	Rank() int
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (sq *Square) Rank() int {
	return 1
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}
func (r Rectangle) Rank() int {
	return 2
}

func main() {
	sq := &Square{10}
	r := Rectangle{3, 5}
	shapers := []Shaper{sq, r}
	for n, _ := range shapers {
		fmt.Println(shapers[n])
		fmt.Println(shapers[n].Area())
	}

	//
	toggen := []TopologicalGenus{sq, r}
	for n, _ := range toggen {
		fmt.Println(toggen[n])
		fmt.Println(toggen[n].Rank())
	}

}
