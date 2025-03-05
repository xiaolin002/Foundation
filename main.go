package main

import (
	"fmt"
	"time"
)

// 函数签名
//func Add() func(b int) int
//func Adder1(a int) func(b int) int

func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder1(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	start := time.Now()
	p2 := Add()
	fmt.Printf("Call Add2 eleven_1 3 gives: %v\n", p2(3))
	Adder2 := Adder1(2)
	fmt.Printf("The result is: %v\n", Adder2(3))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}
