package main

import "fmt"

var i = 5
var str = "abc"
var srtRune = []rune("中国")

type Car struct{}

type Person struct {
	name string
	age  int
}

func main() {
	var val any
	val = i
	fmt.Printf("val has the value：%v\n", val)
	val = str
	fmt.Printf("val has the value：%v\n", val)
	val = srtRune
	fmt.Printf("val has the value：%#v\n", val)
	val = new(Car)
	fmt.Printf("val has the value：%v\n", val)
	val = Person{name: "zhang3", age: 90}
	fmt.Printf("val has the value：%v\n", val)

	switch t := val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case []rune:
		fmt.Println("[]rune")
	case *Car:
		fmt.Println("*Car")
	case Person:
		fmt.Println("Person")
		fmt.Println(t.name)
	case bool:
		fmt.Printf("Type boolean  %T\n", t)
	default:
		fmt.Println("unknown type")
	}

}
