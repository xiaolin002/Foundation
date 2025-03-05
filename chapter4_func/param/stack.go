package main

import "fmt"

/**
 * @Description
 * @Date 2025/1/12 0:01
 **/

func Add(a, b int) {
	fmt.Printf("The sum of %d + %d = %d\n", a, b, a+b)

}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func main() {
	callback(2, Add)
	fv := func(s string) {
		fmt.Println(s)
	}
	fv("hello world")
	fmt.Printf("The type of fv is %T\n", fv)
}
