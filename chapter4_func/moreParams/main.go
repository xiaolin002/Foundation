package main

import "fmt"

func main() {

	num := minNum(1, 3, 2, 0)
	fmt.Printf("The minnum is %d\n", num)

	// 切片也可以用arr...  传递
	a := []int{7, 3, 0, 5}
	x := minNum(a...)
	fmt.Printf("the mininum is: %d\n", x)

}

func minNum(a ...int) int {
	if len(a) == 0 {
		return 0

	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min

}
