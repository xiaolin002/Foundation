package main

import "fmt"

func func1() (i int) {
	defer func() {
		i++
		fmt.Println("defer2,", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1,", i)
	}()
	return 0
}

func main() {
	fmt.Println(func1())
}
